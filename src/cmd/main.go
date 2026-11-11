// Command gnodeb is the gNodeB pseudo-call of kronos-5g-emulator. The
// program is package app; main only connects it to the process.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"kronos/app"
)

// version is the version shown by "gnodeb version" (it can be set when
// building: go build -ldflags "-X main.version=1.0.0").
var version = "0.5.0-dev"

// main runs app.Run, cancelling it on SIGINT or SIGTERM, and exits with its code.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	code := app.Run(ctx, version, os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	stop()
	os.Exit(code)
}
