// Constants of package webconsole.

package webconsole

import "time"

// The HTTP server.
const (
	readHeaderTimeout = 10 * time.Second // time a client has to send the request headers
	shutdownTimeout   = 2 * time.Second  // time other requests get to finish on shutdown
	maxBodySize       = 1 << 20          // the largest JSON request body, in octets
)

// List queries (?limit of /ues and /trace, ?tail of /events).
const (
	defaultListLimit = 1000  // items returned when ?limit is not given
	maxListLimit     = 10000 // larger ?limit and ?tail values are cut to this
)

// The event stream (Server-Sent Events).
const (
	eventBufferSize   = 4096             // live events queued for a slow client before it loses some
	sseRetryMillis    = 2000             // reconnect delay told to the browser, in ms
	keepAliveInterval = 15 * time.Second // a comment line keeps an idle connection open
	maxDrainEvents    = 256              // queued events sent with one flush at most
)
