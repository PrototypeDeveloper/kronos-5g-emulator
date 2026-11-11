package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"kronos/models"
)

// executor runs one command line (implemented by *engine.Engine).
type executor interface {
	Exec(line string) (string, error)
}

// line prints one line; on a terminal it clears the prompt first and redraws it after.
func (c *console) line(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.interactive {
		fmt.Fprint(c.out, "\r\033[K")
	}
	fmt.Fprintln(c.out, s)
	if c.interactive {
		fmt.Fprint(c.out, c.prompt)
		c.shown = true
	}
}

// showPrompt prints the prompt unless a line has just redrawn it.
func (c *console) showPrompt() {
	if !c.interactive {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.shown {
		fmt.Fprint(c.out, c.prompt)
		c.shown = true
	}
}

// entered notes that the user ended a line: the terminal has moved on
// from the prompt.
func (c *console) entered() {
	c.mu.Lock()
	c.shown = false
	c.mu.Unlock()
}

// readCommands executes lines from in until EOF.
func (c *console) readCommands(in io.Reader, x executor) {
	sc := bufio.NewScanner(in)
	c.showPrompt()
	for sc.Scan() {
		c.entered()
		out, err := x.Exec(sc.Text())
		switch {
		case err != nil:
			c.line("error: " + err.Error())
		case out != "":
			for _, l := range strings.Split(out, "\n") {
				c.line(l)
			}
		default:
			c.showPrompt()
		}
	}
}

// formatEvent renders an event as one CUI line, e.g.
//
//	12:00:00.101 node            -> NGSetupRequest                       (NGSetupRequest)
//	12:00:00.123 node            <- NGSetupResponse                      [node]
func formatEvent(e models.Event) string {
	head := e.Time.Format("15:04:05.000") + " " + pad(e.Source, eventSourceWidth)
	switch e.Kind {
	case models.KindMessage:
		arrow := "->"
		if e.Dir == "rx" {
			arrow = "<-"
		}
		label := e.NGAP
		if label == "" {
			label = "?"
		}
		if e.NAS != "" {
			label += ":" + e.NAS
		}
		s := head + " " + arrow + " " + pad(label, eventLabelWidth)
		switch {
		case e.Dir == "tx" && e.Signal != "":
			s += " (" + e.Signal + ")"
		case e.Handler != "":
			s += " [" + e.Handler + "]"
		}
		if e.Text != "" {
			s += " " + e.Text
		}
		return strings.TrimRight(s, " ")
	case models.KindError:
		if e.Handler != "" {
			return head + " error: [" + e.Handler + "] " + e.Text
		}
		return head + " error: " + e.Text
	case models.KindWarn:
		return head + " warning: " + e.Text
	default:
		return head + " " + e.Text
	}
}

// pad right-pads s with spaces to n characters.
func pad(s string, n int) string {
	if w := utf8.RuneCountInString(s); w < n {
		return s + strings.Repeat(" ", n-w)
	}
	return s
}
