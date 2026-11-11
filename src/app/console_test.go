package app

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

type execFunc func(string) (string, error)

func (f execFunc) Exec(line string) (string, error) { return f(line) }

// The prompt is printed once: after the lines printed at start, after an
// empty line, after output and after an error.
func TestConsolePrompt(t *testing.T) {
	var out bytes.Buffer
	c := &console{out: &out, interactive: true, prompt: "gnodeb> "}
	c.line("started")
	c.line("web console: http://127.0.0.1:8080/")
	x := execFunc(func(line string) (string, error) {
		switch line {
		case "help":
			return "linkup\nlinkdown", nil
		case "bad":
			return "", errors.New("unknown command")
		}
		return "", nil
	})
	c.readCommands(strings.NewReader("\nhelp\nbad\n"), x)
	got := strings.ReplaceAll(out.String(), "\r\033[K", "|")
	want := "|started\ngnodeb> |web console: http://127.0.0.1:8080/\ngnodeb> " + // start
		"gnodeb> " + // the empty line
		"|linkup\ngnodeb> |linkdown\ngnodeb> " + // help
		"|error: unknown command\ngnodeb> " // bad
	// (The terminal echoes the user's Enter, so the prompt after the empty
	// line starts a line of its own.)
	if got != want {
		t.Fatalf("got  %q\nwant %q", got, want)
	}
}

// Not on a terminal, there is no prompt.
func TestConsoleNoPrompt(t *testing.T) {
	var out bytes.Buffer
	c := &console{out: &out, prompt: "gnodeb> "}
	c.line("started")
	c.readCommands(strings.NewReader("\n"), execFunc(func(string) (string, error) { return "", nil }))
	if out.String() != "started\n" {
		t.Fatalf("%q", out.String())
	}
}
