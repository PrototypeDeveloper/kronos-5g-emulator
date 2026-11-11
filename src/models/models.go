// Package models holds the data passed between layers: the event stream that
// feeds the trace file, the CUI and the web console, and status
// snapshots. It depends on nothing else in kronos.
package models

import (
	"net/netip"
	"time"
)

// The event stream.

// Event is one entry of the event stream.
type Event struct {
	Seq     uint64    `json:"seq"`
	Time    time.Time `json:"time"`
	Kind    string    `json:"kind"`
	Source  string    `json:"src"`
	Dir     string    `json:"dir,omitempty"` // "tx" or "rx"
	NGAP    string    `json:"ngap,omitempty"`
	NAS     string    `json:"nas,omitempty"`
	Signal  string    `json:"signal,omitempty"`  // the signal.txt name of a sent PDU
	Handler string    `json:"handler,omitempty"` // block whose handler ran ("node", a state), or "unhandled"
	State   string    `json:"state,omitempty"`
	Stream  uint16    `json:"stream,omitempty"`
	Len     int       `json:"len,omitempty"`
	Hex     string    `json:"hex,omitempty"`
	// Plain is the plaintext (hex) of each NAS-PDU of the message, in the
	// order the PDU carries them, "" where the NAS-PDU is in clear. It lets
	// the trace and the web console show ciphered NAS.
	Plain []string `json:"plain,omitempty"`
	Text  string   `json:"text,omitempty"`
}

// Status snapshots of the node and of the subscribers.

// NodeStatus is a snapshot of the node and its peer.
type NodeStatus struct {
	Role       string   `json:"role"`
	Local      string   `json:"local,omitempty"`
	Remote     string   `json:"remote,omitempty"`
	Connected  bool     `json:"connected"`
	OutStreams uint16   `json:"outStreams,omitempty"`
	InStreams  uint16   `json:"inStreams,omitempty"`
	Setup      string   `json:"setup"`
	PeerName   string   `json:"peerName,omitempty"` // AMF name (gNB role)
	PeerInfo   []string `json:"peerInfo,omitempty"` // served GUAMIs, supported PLMNs and slices
	Cause      string   `json:"cause,omitempty"`    // why NG Setup failed
}

// UEStatus is a snapshot of one subscriber.
type UEStatus struct {
	IMSI     string            `json:"imsi"`
	State    string            `json:"state"`             // scenario state
	LastMsg  string            `json:"lastMsg,omitempty"` // last received NGAP[:NAS]
	RM       string            `json:"rm"`
	CM       string            `json:"cm"`
	GUTI     string            `json:"guti,omitempty"` // 5GS mobile identity value (hex)
	RANID    *int64            `json:"ranUeNgapId,omitempty"`
	AMFID    *int64            `json:"amfUeNgapId,omitempty"`
	Security *NASSecurity      `json:"security,omitempty"`
	Sessions []SessionStatus   `json:"sessions,omitempty"`
	Vars     map[string]string `json:"vars,omitempty"` // private variables (literals)
}

// NASSecurity is the current NAS security context of a subscriber.
type NASSecurity struct {
	NgKSI     uint8  `json:"ngKsi"`
	Integrity string `json:"integrity"` // NIA0..NIA3
	Ciphering string `json:"ciphering"` // NEA0..NEA3
	ULCount   uint32 `json:"ulCount"`
	DLCount   uint32 `json:"dlCount"`
}

// SessionStatus is one PDU session of a subscriber.
type SessionStatus struct {
	ID     uint8  `json:"id"`
	UEIP   string `json:"ueIp,omitempty"`
	UPF    string `json:"upf,omitempty"`
	ULTEID uint32 `json:"ulTeid,omitempty"`
	DLTEID uint32 `json:"dlTeid,omitempty"`
	QFIs   []int  `json:"qfi,omitempty"` // ints: a []uint8 would be base64 in JSON
	Active bool   `json:"active"`        // the N3 tunnel is up
}

// Counters and the dashboard.

// Counters count what went through the event stream.
type Counters struct {
	TX        uint64                   `json:"tx"`
	RX        uint64                   `json:"rx"`
	Errors    uint64                   `json:"errors"`
	Warnings  uint64                   `json:"warnings"`
	Unhandled uint64                   `json:"unhandled"`
	Dropped   uint64                   `json:"dropped"`  // events lost by slow subscribers
	Messages  map[string]*MessageCount `json:"messages"` // by NGAP name
}

// MessageCount counts the PDUs of one NGAP message type.
type MessageCount struct {
	TX uint64 `json:"tx"`
	RX uint64 `json:"rx"`
}

// Overview is the dashboard of the web console.
type Overview struct {
	Role        string         `json:"role"`
	Config      string         `json:"config"`
	Started     time.Time      `json:"started"`
	Snap        bool           `json:"snap"`
	Node        NodeStatus     `json:"node"`
	Subscribers int            `json:"subscribers"`
	States      map[string]int `json:"states"` // subscribers per scenario state
	RM          map[string]int `json:"rm"`
	CM          map[string]int `json:"cm"`
	Sessions    int            `json:"sessions"` // PDU sessions with the N3 tunnel up
	Counters    Counters       `json:"counters"`
}

// Signals, the scenario and decoded messages as the web console shows them.

// SignalInfo describes one template of signal.txt.
type SignalInfo struct {
	Name  string `json:"name"`
	Scope string `json:"scope"` // "UE", "node", or "?" when it could not be decoded
	Mode  string `json:"mode"`
	NGAP  string `json:"ngap,omitempty"`
	Label string `json:"label"` // NGAP[:NAS] as receive() triggers name it
	Line  int    `json:"line"`
	Hex   string `json:"hex"`
}

// ScenarioInfo is the scenario as the web console shows it.
type ScenarioInfo struct {
	Path   string      `json:"path"`
	Vars   []VarValue  `json:"vars"`
	Blocks []BlockInfo `json:"blocks"`
}

// BlockInfo is one node, any or state block.
type BlockInfo struct {
	Kind     string        `json:"kind"` // "node", "any" or "state"
	Name     string        `json:"name"`
	Line     int           `json:"line"`
	Entry    []string      `json:"entry,omitempty"`
	Handlers []HandlerInfo `json:"handlers,omitempty"`
}

// HandlerInfo is one receive() handler.
type HandlerInfo struct {
	Trigger string   `json:"trigger"`
	Line    int      `json:"line"`
	Calls   []string `json:"calls"`
}

// VarValue is a scenario variable: the current value of a public one, the
// initial value of a private one.
type VarValue struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Public bool   `json:"public"`
	Value  string `json:"value"`
}

// TreeNode is one line of a decoded message: an IE and its value, with the
// IEs inside it (the sequence view of the web console).
type TreeNode struct {
	Name  string      `json:"n"`
	Value string      `json:"v,omitempty"`
	Kids  []*TreeNode `json:"c,omitempty"`
}

// The signal editor.

// EditNode is a node of a signal in the signal editor: a field with its
// value, what it may hold, and what can be done to it.
type EditNode struct {
	Name  string `json:"n"`
	Path  string `json:"p"`           // where set / add / choose / hex act
	Kind  string `json:"k"`           // struct, choice, list, int, octets, bits, ip, plmn, string, bool, container, ext
	Value string `json:"v,omitempty"` // the value as it is edited (hex for octets and containers)
	Hint  string `json:"h,omitempty"` // a reading of the value (MCC/MNC, hex of a number, size)
	Bits  int    `json:"bits,omitempty"`
	// Min and Max bound the value (int) or the size (octets, bits,
	// string, list); Ext allows more (an extensible constraint).
	Min  *int64      `json:"min,omitempty"`
	Max  *int64      `json:"max,omitempty"`
	Ext  bool        `json:"ext,omitempty"`
	Del  string      `json:"del,omitempty"`  // the path that remove takes: an optional field or a list item
	Add  []string    `json:"add,omitempty"`  // struct: the optional fields that are absent
	Alts []string    `json:"alts,omitempty"` // choice: the alternatives
	Alt  string      `json:"alt,omitempty"`  // choice: the one chosen
	Auto bool        `json:"auto,omitempty"` // mapped: replaced when the signal is sent (still editable)
	Note string      `json:"note,omitempty"`
	Kids []*EditNode `json:"c,omitempty"`
}

// DiagText is a problem found in a signal or in signal.txt.
type DiagText struct {
	Sev  string `json:"sev"` // error, warning, info
	Line int    `json:"line,omitempty"`
	Msg  string `json:"msg"`
}

// SignalFileEntry is a signal as signal.txt has it now.
type SignalFileEntry struct {
	Name   string `json:"name"`
	Mode   string `json:"mode"`
	Line   int    `json:"line"`
	Hex    string `json:"hex"`
	Label  string `json:"label"` // NGAP:NAS
	Loaded bool   `json:"loaded"`
}

// SignalFile is signal.txt as the editor reads it.
type SignalFile struct {
	Path    string            `json:"path"`
	Signals []SignalFileEntry `json:"signals"`
	Diags   []DiagText        `json:"diags"`
	// Pending reports that the file differs from the signals in use
	// (reload uses it).
	Pending bool `json:"pending"`
}

// SignalEdit is a signal in the editor: its octets, its tree and its checks.
type SignalEdit struct {
	Hex   string     `json:"hex"`
	Tree  *EditNode  `json:"tree"`
	Diags []DiagText `json:"diags"`
}

// Arguments of the subscriber commands (package command fills them, the engine runs them).

// PingArgs are the arguments of ping.
type PingArgs struct {
	Dst      string
	Count    int           // 0: until stop
	Interval time.Duration // 0: the default (1 s)
	Size     int           // ICMP payload octets; negative: the default (56)
}

// IperfArgs are the arguments of iperf.
type IperfArgs struct {
	Dst      string
	Port     int
	Rate     uint64 // bit/s; 0: the default
	Duration time.Duration
	Reverse  bool // dl: the server sends
	Length   int  // UDP payload or TCP write octets; 0: the default
	TCP      bool // iperftcp: a TCP test through the UE's TUN device
}

// What a role gives the engine (package role and its implementations).

// Built is a PDU ready to send.
type Built struct {
	PDU []byte
	// Plain is the plaintext of each NAS-PDU of the PDU, in the order of
	// scenario.NASPDUs; nil where the NAS-PDU is in clear.
	Plain [][]byte
}

// Inbound is what the built-in processing found in a received PDU.
type Inbound struct {
	NAS     []string // NAS message names, in trigger priority order
	Notes   []string // results worth showing: authentication, security mode, registration
	Errors  []string // problems: MAC failure, undecodable NAS ...
	Release bool     // UE context release: forget the NGAP IDs after the handlers ran
	// TunnelsUp lists the PDU sessions whose user plane became usable;
	// TunnelsDown the DL TEIDs of tunnels that went away.
	TunnelsUp   []Tunnel
	TunnelsDown []uint32
	// PagingKey identifies the UE in Paging (scenario.PagingKey) once the
	// network has assigned a 5G-GUTI.
	PagingKey string
	// Plain is the plaintext of each NAS-PDU, in the order of
	// scenario.NASPDUs; nil where the NAS-PDU was in clear or could not be opened.
	Plain [][]byte
}

// Tunnel is the N3 GTP-U tunnel of one PDU session.
type Tunnel struct {
	SessionID uint8
	UEIP      netip.Addr     // PDU address given by the network
	UPF       netip.AddrPort // where uplink G-PDUs go
	ULTEID    uint32         // TEID the UPF expects
	DLTEID    uint32         // TEID the gNB allocated for downlink
	QFI       uint8          // QoS flow for uplink packets
}

// Signal editing and checks (package scenario).

// EditOp is one change made in the editor.
type EditOp struct {
	Kind  string `json:"kind"`            // set, add, remove, choose, hex
	Path  string `json:"path"`            // the node's path (for remove, its del path)
	Value string `json:"value,omitempty"` // set, hex: the value; add on a struct: the field; choose: the alternative
	Bits  int    `json:"bits,omitempty"`  // set on a BIT STRING: its length in bits (0 keeps it)
}

// NASInfo describes one 5GS NAS message type (TS 24.501 Table 9.7.1 and 9.7.2).
type NASInfo struct {
	Name      string
	EPD       uint8
	Type      uint8
	Supported bool // libnas decodes the body; otherwise only the name is known
}

// SignalReference holds the config values that auto-fix writes into templates.
// SignalSet.Compare (package scenario) reports template values that differ from them.
type SignalReference struct {
	MCC, MNC string
	TAC      uint32
	HasTAC   bool
	NCI      uint64 // 36-bit NR Cell Identity
	HasNCI   bool
	GNBID    uint32
	GNBIDLen int
	HasGNBID bool
	Name     string // RANNodeName; "" keeps the template value
}

// ValidateOptions are the config settings that affect validation.
type ValidateOptions struct {
	SnapEnabled bool // base.snap
}

// Parsed packets of the UE IP stack (package traffic).

// IPv4 is a parsed IPv4 packet.
type IPv4 struct {
	Src, Dst netip.Addr
	Proto    uint8
	TTL      uint8
	ID       uint16
	Frag     bool   // a fragment (more fragments, or a non-zero offset)
	Payload  []byte // after the header; aliases the packet
}

// UDP is a parsed UDP datagram.
type UDP struct {
	Src, Dst netip.AddrPort
	Payload  []byte // aliases the packet
}

// Echo is a parsed ICMP echo request or reply.
type Echo struct {
	Type    uint8 // 8 request, 0 reply
	ID, Seq uint16
	Data    []byte // aliases the packet
}

// ping options (package traffic/ping).

// PingOptions are the options of one ping run.
type PingOptions struct {
	Dst      netip.Addr
	Count    int           // requests to send; 0 sends until ctx ends
	Interval time.Duration // between requests (ping.DefaultInterval if 0)
	Size     int           // ICMP payload octets (ping.DefaultSize if negative)
}

// SCTP (package transport/sctp).

// SCTPInfo describes a received message.
type SCTPInfo struct {
	Stream uint16
	PPID   uint32
}

// Web console (package webconsole).

// WebOptions are the settings of the console.
type WebOptions struct {
	Version string // shown in the GUI
}
