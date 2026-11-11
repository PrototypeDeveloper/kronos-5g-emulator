// Constants of package app.

package app

// The command line.

const progName = "gnodeb"

const usage = `Usage:
  gnodeb [-c config.yaml] [--no-cui]   start the pseudo-call (linkup connects N2); gnodeb.web of the config also opens the web console
  gnodeb [-c config.yaml] --exec FILE  run the commands of FILE in order and exit (batch; exit code 1 on a failure)
  gnodeb check [-c config.yaml] [-v]   check config, signal and scenario
  gnodeb version                       print the version

Options:
  -c FILE      path of config.yaml (default config/config.yaml); scenario.txt and signal.txt are read from the same directory
  -v           with check, also print what was understood from the signals and the scenario
  --no-cui     run without the CUI (headless; stop with Ctrl-C or SIGTERM); without a web console, linkup runs by itself
  --no-web     do not start the web console
  --exec FILE  run the commands one line at a time and exit (batch); @FILE in the CUI does the same
`

// exitUsage is the exit code for invalid arguments or an unknown subcommand.
const exitUsage = 2

// CUI event lines.
const (
	// eventSourceWidth is the column width of the event source.
	eventSourceWidth = 15
	// eventLabelWidth is the column width of the message label.
	eventLabelWidth = 44
)

// Event subscriptions.
const (
	// consoleEventBuffer is the event buffer of the CUI and log subscriber.
	consoleEventBuffer = 4096
	// traceEventBuffer is the event buffer of the trace writer.
	traceEventBuffer = 8192
)

// The trace file.
const (
	// traceDirPerm is the permission of the trace directory.
	traceDirPerm = 0o755
	// traceFilePerm is the permission of the trace file.
	traceFilePerm = 0o644
)
