// Types of the engine.

package engine

import (
	"context"
	"net"
	"net/netip"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"kronos/command"
	"kronos/models"
	"kronos/scenario"
	"kronos/simulator/role"
	"kronos/state"
	"kronos/traffic"
	"kronos/transport/peer"
	"kronos/transport/sctp"
)

// The engine and the state it shares between actors.

// Engine runs one pseudo-call.
type Engine struct {
	cfg          *state.Config
	sigs         atomic.Pointer[scenario.SignalSet] // replaced by reload
	fileMu       sync.Mutex                         // signal.txt writes and reloads
	scen         *scenario.File
	role         role.Role
	bus          *Bus
	vars         *varStore
	snap         bool
	nodeHandlers map[string]*scenario.Handler // by canonical NGAP name
	subs         []*subscriber
	nodeMapOff   atomic.Uint32 // models.MapOff of the node's signals (kept across linkdown)
	controls     *controls     // the turns of control()
	peer         *peer.Client  // sync_send; nil without gnodeb.http
	routes       *routes
	uplane       *uplane
	openTUN      tunOpener                       // nil without CAP_NET_ADMIN
	ctx          atomic.Pointer[context.Context] // set by Start (runCtx reads it); ends background work such as ping
	started      time.Time

	mu       sync.Mutex
	running  bool
	linkBusy bool  // a linkup is in progress
	node     *node // the current association; nil while the link is down
	stopped  chan struct{}
}

// consoleEnv runs the commands that act on the pseudo-call as a whole
// (help, linkup, linkdown) and collects their output.
type consoleEnv struct {
	e   *Engine
	out *strings.Builder
}

// controls paces the control(n) statements of the scenario: the turn last
// given at each (by its position), shared by all subscribers.
type controls struct {
	mu   sync.Mutex
	last map[string]time.Time
	cut  chan struct{} // closed when the link goes down, which ends the waits
}

// varStore holds the public variables, shared by the whole pseudo-call. Each
// access is atomic, and so is each set (assign).
type varStore struct {
	mu       sync.Mutex
	m        map[string]command.Value
	assignMu sync.Mutex // one set at a time: its value is computed and stored under it
}

// routes maps the UE NGAP IDs of current UE associations to subscribers,
// and the 5G-S-TMSI of registered UEs (kept in CM-IDLE) for Paging.
type routes struct {
	mu     sync.Mutex
	ran    map[int64]*subscriber
	amf    map[int64]*subscriber
	paging map[string]*subscriber
	// awaitHO are the subscribers that took a handover context over
	// (sync_receive(HO)) and wait for the AMF's HandoverRequest, oldest first.
	awaitHO []hoWait
}

// hoWait is a subscriber waiting for its HandoverRequest, with the uplink
// TEIDs of its sessions, by which the request is matched to it.
type hoWait struct {
	s       *subscriber
	ulTEIDs []uint32
}

// The event bus.

// Bus fans events out to subscribers (CUI, trace file, web console) and keeps
// recent history. Publishing never blocks: a subscriber that falls behind
// loses events and the loss is counted.
type Bus struct {
	mu      sync.Mutex
	seq     uint64
	subs    map[int]chan models.Event
	next    int
	ring    []models.Event
	head    int
	full    bool
	dropped uint64
	count   models.Counters
}

// The node actor and the messages of the actors' inboxes.

// node is the actor of one SCTP association. Received PDUs,
// the node block of the scenario and node commands all run on its
// goroutine. UE-associated PDUs are handed to the subscriber actors.
type node struct {
	e           *Engine
	conn        *sctp.Conn
	inbox       chan any
	done        chan struct{}
	lastMsg     string
	subsStarted bool
	outStreams  uint16   // fixed for the association; read by subscriber actors
	timers      timers   // the node's scenario timers (public variables)
	rx          received // the last received message ($ngap)

	mu     sync.Mutex // guards status for snapshots taken by other goroutines
	status models.NodeStatus
}

// rxMsg is a PDU received on the association.
type rxMsg struct {
	b    []byte
	info models.SCTPInfo
}

// lostMsg tells the actor that the association has ended.
type lostMsg struct{ err error }

// execMsg runs a command on an actor; the result goes to reply.
type execMsg struct {
	call  command.Call
	reply chan<- execResult
}

// execResult is the outcome of an execMsg; imsi names the actor that ran it.
type execResult struct {
	imsi string
	err  error
}

// startMsg makes the actor run the entry of its start state (NG Setup has succeeded).
type startMsg struct{}

// resetMsg makes the actor return to its initial state (the link went down).
type resetMsg struct{}

// releaseMsg ends the subscriber's UE association (NG Reset); its NAS state stays.
type releaseMsg struct {
	why string
}

// ueRx is a UE-associated PDU routed to this subscriber, already decoded.
type ueRx struct {
	msg  any
	raw  []byte
	info models.SCTPInfo
}

// The subscriber actor, its traffic runs and its TUN device.

// subscriber is the actor of one subscriber: the scenario state
// machine, the private variables and the role's protocol context. Everything
// runs on its goroutine, so none of it needs locking.
type subscriber struct {
	e     *Engine
	index int
	imsi  string
	inbox chan any
	ue    role.UE

	state   string
	vars    map[string]command.Value
	mapOff  atomic.Uint32 // models.MapOff: the mapping items off for this subscriber (kept across linkdown)
	rx      received      // the last received message ($ngap, $nas)
	lastMsg string
	moved   string                  // the last move() target, whose entry runs after the handler
	tunnels map[uint8]models.Tunnel // active PDU session tunnels by session ID
	stack   *traffic.Stack          // the UE's IP endpoints (ping, iperf); safe for concurrent use
	runs    trafficRuns             // the pings and iperfs running
	timers  timers                  // the running scenario timers (timer(t))
	tun     atomic.Pointer[ueTUN]   // the host TUN device; the N3 reader writes to it

	// stat is the snapshot the web console reads without stopping the actor;
	// the actor refreshes it after each message.
	statMu sync.Mutex
	stat   models.UEStatus
}

// trafficRun is one running ping or iperf.
type trafficRun struct {
	kind   string // "ping" or "iperf"
	cancel context.CancelFunc
}

// trafficRuns holds a subscriber's runs by ID; a run removes itself from its own goroutine when it ends.
type trafficRuns struct {
	mu   sync.Mutex
	next int
	runs map[int]trafficRun
}

// ueTUN is the device of a subscriber and the PDU session it carries.
type ueTUN struct {
	dev     tunDevice
	session uint8
	ueIP    netip.Addr
	dropped atomic.Uint64 // uplink packets with no tunnel (idle) or not from the UE
}

// The N3 user plane.

// uplane is the N3 user plane: one kernel UDP socket on
// gtpu.src. Downlink G-PDUs go by TEID to the subscriber's IP stack (ping,
// iperf), ICMP echo requests to a UE are answered, and GTP-U Echo Requests
// too. The socket is opened with the first tunnel.
type uplane struct {
	e         *Engine
	mu        sync.Mutex
	conn      *net.UDPConn
	byDL      map[uint32]dlRoute
	bySession map[sessionKey]models.Tunnel
}

// dlRoute is where a downlink TEID leads: the subscriber and its tunnel.
type dlRoute struct {
	s *subscriber
	t models.Tunnel
}

// sessionKey names a PDU session of a subscriber.
type sessionKey struct {
	s  *subscriber
	id uint8
}

// Scenario timers (timer.go).

// timerRun is a running timer; gen tells a late expiry of an earlier run apart.
type timerRun struct {
	t   *time.Timer
	gen uint64
}

// syncMsg hands an actor a message of another pseudo-call (sync_receive).
type syncMsg struct {
	m models.SyncMessage
}

// timerMsg tells an actor that the timer of a variable ran out.
type timerMsg struct {
	name string
	gen  uint64
}

// timers are the running timers of an actor (a subscriber or the node).
type timers struct {
	runs map[string]*timerRun
	gen  uint64
}

// received is the message that started the handler (or came last), which
// $ngap.X and $nas.X read; its NAS messages are decoded on first use.
type received struct {
	ngap    any      // the decoded NGAP message
	plain   [][]byte // the plain text of its NAS-PDUs that were protected (nil: as sent)
	decoded []any    // its NAS messages, once decoded
	done    bool     // decoded is set
}
