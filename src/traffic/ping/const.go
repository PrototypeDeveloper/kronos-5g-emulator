// Constants of package ping.

package ping

import "time"

// Defaults and limits.
const (
	DefaultInterval = time.Second
	DefaultSize     = 56
	MaxSize         = 65000
	// MinInterval keeps a runaway command from flooding the tunnel.
	MinInterval = time.Millisecond
	// QuietBelow: with a shorter interval, each reply is not reported (a
	// progress line comes every second instead).
	QuietBelow = 200 * time.Millisecond
	timeout    = 2 * time.Second
	// stopGrace is how long a stopped ping waits for the replies in flight.
	stopGrace = 500 * time.Millisecond
)

// The loop of a run.
const (
	checkInterval = 100 * time.Millisecond // how often the requests in flight are checked for a timeout
	replyQueueLen = 1024                   // replies queued for the loop
)

// ICMP types.
const (
	icmpEchoRequest = 8
	icmpEchoReply   = 0
)
