// Types of package ping.

package ping

import (
	"time"
)

// Stats is the outcome of a run.
type Stats struct {
	Sent, Received  int
	SendErrors      int
	Min, Max        time.Duration
	sum, sumSquares float64 // of the RTTs in milliseconds
	Elapsed         time.Duration
}

// reply is an echo reply received: its sequence number and arrival time.
type reply struct {
	seq uint16
	at  time.Time
}
