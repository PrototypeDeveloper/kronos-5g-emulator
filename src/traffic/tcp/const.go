// Constants of package tcp.

package tcp

import "time"

// Header flags.
const (
	finFlag = 0x01
	synFlag = 0x02
	rstFlag = 0x04
	pshFlag = 0x08
	ackFlag = 0x10
)

// MSS is the largest segment payload sent and announced, so the inner
// packets stay within the GTP-U path MTU.
const MSS = 1360

// Retransmission and window limits.
const (
	rtoInitial = 300 * time.Millisecond
	rtoMax     = 3 * time.Second
	maxRetries = 8
	sendWindow = 8 * MSS
	recvWindow = 65535
	finWait    = 2 * time.Second // how long a closed connection waits for the peer's FIN
)

// Phases of a connection.
const (
	synSent state = iota
	synRcvd
	established
)

// Timer and queues.
const (
	timerTick      = 20 * time.Millisecond // how often a connection checks its retransmission timer
	inQueueLen     = 256                   // received segments queued for a connection's goroutine
	acceptQueueLen = 16                    // established connections waiting for Accept
)
