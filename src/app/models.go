// Structs of package app.

package app

import (
	"io"
	"sync"
)

// options holds the parsed command-line flags.
type options struct {
	version string
	config  string
	verbose bool
	noCUI   bool
	noWeb   bool
	exec    string
}

// console is the CUI: one command per line, events printed as
// they happen. On a terminal the prompt is redrawn after asynchronous lines.
type console struct {
	mu          sync.Mutex
	out         io.Writer
	interactive bool
	prompt      string
	shown       bool // the prompt is on the last line of the terminal
}
