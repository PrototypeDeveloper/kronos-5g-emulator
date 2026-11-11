// Types of package tcp.

package tcp

import (
	"net/netip"
	"sync"
	"time"

	"kronos/traffic"
)

// segment is a received TCP segment.
type segment struct {
	src      netip.AddrPort
	seq, ack uint32
	flags    uint8
	window   uint16
	payload  []byte
}

// Conn is one TCP connection. Its methods are safe for concurrent use.
type Conn struct {
	st      *traffic.Stack
	link    traffic.Link
	local   netip.AddrPort
	remote  netip.AddrPort
	in      chan segment
	done    chan struct{} // closed when the connection has ended
	release func()        // frees the endpoint
	// onEstablished runs (with mu held) when a passive open completes.
	onEstablished func()

	mu          sync.Mutex
	changed     chan struct{} // closed on every change (wakes the waiters)
	st8         state
	iss         uint32
	sndUna      uint32 // oldest unacknowledged sequence number
	sndNxt      uint32 // next sequence number to send
	sndBuf      []byte // data from sndUna on: in flight, then not sent yet
	peerWnd     uint32
	finQueued   bool // Close was called: a FIN follows the data
	finSent     bool
	finAcked    bool
	finAckedAt  time.Time
	rcvNxt      uint32
	rcvBuf      []byte
	rcvFIN      bool
	err         error
	rto         time.Duration
	retries     int
	lastSend    time.Time // when the oldest outstanding segment was (re)sent
	readTimeout time.Time
}

// Listener accepts connections on one local port.
type Listener struct {
	st     *traffic.Stack
	link   traffic.Link
	port   uint16
	accept chan *Conn
	mu     sync.Mutex
	conns  map[netip.AddrPort]*Conn
	closed chan struct{}
	once   sync.Once
}
