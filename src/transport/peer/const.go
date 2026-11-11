// Constants of package peer.

package peer

import "time"

// SyncPath is where a pseudo-call takes the messages of the others (POST, JSON).
const SyncPath = "/kronos/v1/sync"

// Limits of the exchange.
const (
	sendTimeout       = 5 * time.Second  // a sync_send waits this long for the answer
	readHeaderTimeout = 10 * time.Second // the server reads the request header within this
	maxBody           = 1 << 20          // octets of a request or answer body
)
