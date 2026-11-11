// Data types of package webconsole.

package webconsole

import (
	"net/http"
	"sync"

	"kronos/models"
)

// The server.

// Server serves the GUI and the API.
type Server struct {
	b    Backend
	opts models.WebOptions
	mux  *http.ServeMux

	stopOnce sync.Once
	stop     chan struct{} // closed on shutdown: ends the event streams
}

// API replies.

// statusReply is the overview plus the version of the binary.
type statusReply struct {
	models.Overview
	Version string `json:"version"`
}

// treeReply is a signal or an event with its PDU decoded as a tree.
type treeReply struct {
	Signal *models.SignalInfo `json:"signal,omitempty"`
	Event  *models.Event      `json:"event,omitempty"`
	Tree   *models.TreeNode   `json:"tree,omitempty"`
	Error  string             `json:"error,omitempty"` // the PDU could not be decoded (fully)
}

// commandInfo is a registered command as the GUI builds its forms from it.
type commandInfo struct {
	Name  string    `json:"name"`
	Args  []argInfo `json:"args"`
	Scope []string  `json:"scope"` // node, ue, console, signal (send: decided by the signal)
	Where []string  `json:"where"` // scenario, interactive
	Help  string    `json:"help"`
}

// argInfo is one parameter of a command in commandInfo.
type argInfo struct {
	Name     string   `json:"name"`
	Kind     string   `json:"kind"` // signal, state, format, value, command, rate, choice
	Optional bool     `json:"optional,omitempty"`
	Variadic bool     `json:"variadic,omitempty"`
	Choices  []string `json:"choices,omitempty"` // choice: the words allowed
	Default  string   `json:"default,omitempty"` // what an omitted optional argument means
}

// commandReply is the reply of POST /api/v1/commands.
type commandReply struct {
	Line   string `json:"line"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

// API requests.

// signalRequest is the body of the signal editor requests.
type signalRequest struct {
	Name string         `json:"name"`
	Hex  string         `json:"hex"`
	Mode string         `json:"mode"`
	Op   *models.EditOp `json:"op,omitempty"`
}
