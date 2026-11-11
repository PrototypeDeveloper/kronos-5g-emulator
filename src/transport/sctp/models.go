// Types of package sctp.

package sctp

import (
	"net/netip"
	"os"
	"sync/atomic"
	"syscall"
	"time"
)

// Options tune a socket before the association is set up.
type Options struct {
	OutStreams uint16        // outbound streams to request (default 16)
	InStreams  uint16        // inbound streams to accept (default 16)
	Timeout    time.Duration // Dial: time allowed to set up the association (default 10s)
}

// Conn is one SCTP association. Send may be called from several goroutines;
// Recv from one goroutine at a time.
type Conn struct {
	f       *os.File
	rc      syscall.RawConn
	local   netip.AddrPort
	remote  netip.AddrPort
	in, out uint16
	buf     []byte
	oob     []byte
	closed  atomic.Bool
}

// Listener accepts associations (for the AMF role, not implemented yet; also used by tests).
type Listener struct {
	f      *os.File
	rc     syscall.RawConn
	addr   netip.AddrPort
	opt    Options
	closed atomic.Bool
}
