// Package app is the gnodeb program: the arguments and
// subcommands, the start of a pseudo-call (engine, web console, CUI, trace
// and log), the batch mode and the static check. cmd/main.go only connects
// it to the process.
package app

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
)

// Run is the gnodeb program: it parses the arguments, runs the subcommand
// and returns the exit code. ctx ends the pseudo-call (Ctrl-C, SIGTERM).
func Run(ctx context.Context, version string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	mode := "run"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		mode, args = args[0], args[1:]
	}
	o := options{version: version}
	fs := flag.NewFlagSet(progName, flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() { fmt.Fprint(stderr, usage) }
	fs.StringVar(&o.config, "c", "config/config.yaml", "")
	fs.BoolVar(&o.verbose, "v", false, "")
	fs.BoolVar(&o.noCUI, "no-cui", false, "")
	fs.BoolVar(&o.noWeb, "no-web", false, "")
	fs.StringVar(&o.exec, "exec", "", "")
	if err := fs.Parse(args); err != nil {
		return exitUsage
	}
	// Accept flags before the subcommand as well: gnodeb -c x.yaml check
	if mode == "run" && fs.NArg() > 0 {
		mode = fs.Arg(0)
		if err := fs.Parse(fs.Args()[1:]); err != nil {
			return exitUsage
		}
	}
	if fs.NArg() > 0 {
		fmt.Fprintf(stderr, "unexpected arguments: %s\n\n%s", strings.Join(fs.Args(), " "), usage)
		return exitUsage
	}
	switch mode {
	case "check":
		return check(o.config, o.verbose, stdout)
	case "version":
		fmt.Fprintln(stdout, progName, version)
		return 0
	case "help":
		fmt.Fprint(stdout, usage)
		return 0
	case "run":
		return start(ctx, o, stdin, stdout, stderr)
	}
	fmt.Fprintf(stderr, "unknown subcommand %s\n\n%s", mode, usage)
	return exitUsage
}
