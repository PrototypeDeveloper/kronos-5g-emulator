// Types of package iperf.

package iperf

import (
	"io"
	"net/netip"
	"sync"
	"time"
)

// UDP tests.

// Options of one test.
type Options struct {
	Server   netip.AddrPort
	Rate     uint64        // bit/s (models.IperfDefaultRate if 0)
	Duration time.Duration // whole seconds (models.IperfDefaultDuration if 0)
	Reverse  bool          // the server sends: a downlink test (iperf3 -R)
	Length   int           // UDP payload octets (models.IperfDefaultLength if 0)
}

// Result is the outcome of a test, from both ends.
type Result struct {
	Reverse     bool
	Seconds     float64
	SentBytes   uint64
	SentPackets uint64
	RecvBytes   uint64
	Lost        uint64
	Total       uint64  // datagrams the sender counted
	Jitter      float64 // milliseconds, measured by the receiver
	OutOfOrder  uint64  // seen by this end when it receives
}

// TCP tests.

// TCPOptions of one TCP test.
type TCPOptions struct {
	Server   netip.AddrPort
	Local    netip.Addr    // the UE address the sockets bind to
	Rate     uint64        // bit/s; 0: as fast as TCP goes
	Duration time.Duration // whole seconds (models.IperfDefaultDuration if 0)
	Reverse  bool          // the server sends: a downlink test (iperf3 -R)
	Length   int           // octets per write (models.IperfDefaultTCPLength if 0)
}

// TCPResult is the outcome of a TCP test, from both ends.
type TCPResult struct {
	Reverse     bool
	Seconds     float64
	SentBytes   uint64
	RecvBytes   uint64
	Retransmits int64 // of the sending end; -1 when unknown
}

// The control connection and its JSON messages.

// params are the test parameters the client sends (iperf_api.c send_parameters).
type params struct {
	UDP           bool   `json:"udp,omitempty"`
	TCP           bool   `json:"tcp,omitempty"`
	Omit          int    `json:"omit"`
	Time          int    `json:"time"`
	Num           int    `json:"num"`
	BlockCount    int    `json:"blockcount"`
	Parallel      int    `json:"parallel"`
	Reverse       bool   `json:"reverse,omitempty"`
	Len           int    `json:"len"`
	Bandwidth     uint64 `json:"bandwidth"`
	PacingTimer   int    `json:"pacing_timer"`
	ClientVersion string `json:"client_version"`
}

// streamResult is one stream in the results exchanged at the end.
type streamResult struct {
	ID             int     `json:"id"`
	Bytes          float64 `json:"bytes"`
	Retransmits    float64 `json:"retransmits"`
	Jitter         float64 `json:"jitter"` // seconds
	Errors         float64 `json:"errors"`
	OmittedErrors  float64 `json:"omitted_errors"`
	Packets        float64 `json:"packets"`
	OmittedPackets float64 `json:"omitted_packets"`
	StartTime      float64 `json:"start_time"`
	EndTime        float64 `json:"end_time"`
}

// results are what each end sends at EXCHANGE_RESULTS.
type results struct {
	CPUUtilTotal         float64        `json:"cpu_util_total"`
	CPUUtilUser          float64        `json:"cpu_util_user"`
	CPUUtilSystem        float64        `json:"cpu_util_system"`
	SenderHasRetransmits int            `json:"sender_has_retransmits"`
	Streams              []streamResult `json:"streams"`
}

// control is the control connection: the UE's small TCP (UDP tests) or a
// host socket (TCP tests).
type control struct {
	c        io.ReadWriter
	deadline func(time.Time) // of reads
}

// The UDP data stream.

// sender paces datagrams of a stream at a bit rate.
type sender struct {
	rate     uint64
	length   int
	send     func(payload []byte) error
	start    time.Time
	payload  []byte
	pcount   uint32
	bytes    uint64
	packets  uint64
	errors   int
	firstErr error
}

// receiver measures a UDP stream as iperf3 does (iperf_udp_recv): loss from
// gaps in the packet counter, jitter by RFC 1889 from the sender times.
type receiver struct {
	mu          sync.Mutex
	bytes       uint64
	datagrams   uint64
	count       uint64 // highest packet counter seen
	lost        uint64
	outOfOrder  uint64
	jitter      float64 // seconds
	prevTransit float64
	started     bool
}

// receiverSnapshot is a copy of the receiver's counters.
type receiverSnapshot struct {
	bytes, datagrams, count, lost, outOfOrder uint64
	jitter                                    float64
}

// ctlRead is one octet read from the control connection during a TCP test.
type ctlRead struct {
	n   int
	b   byte
	err error
}
