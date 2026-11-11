// Data types of package peer.

package peer

import (
	"net/http"

	"kronos/models"
)

// Server takes the messages of other pseudo-calls and hands them to Deliver.
type Server struct {
	mux     *http.ServeMux
	deliver func(models.SyncMessage) error
}

// Client sends messages to other pseudo-calls from the address of gnodeb.http.src.
type Client struct {
	http *http.Client
	from string // gnodeb.http.src as host:port, told to the receiver
}
