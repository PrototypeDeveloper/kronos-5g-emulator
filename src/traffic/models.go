// Types of package traffic.

package traffic

import (
	"sync"
	"sync/atomic"
)

// Parsed packets.

// The endpoint table.

// endpoint is the key of a handler: a protocol and a port.
type endpoint struct {
	proto uint8
	port  uint16 // UDP/TCP: the local port; ICMP: the echo identifier
}

// Stack is the endpoint table of one UE: which session gets the downlink
// packets of each protocol and port. It is safe for concurrent use.
type Stack struct {
	mu       sync.Mutex
	handlers map[endpoint]Handler
	next     map[uint8]uint16 // last ephemeral port (or echo identifier) given out
	ipID     atomic.Uint32
	unknown  atomic.Uint64
}
