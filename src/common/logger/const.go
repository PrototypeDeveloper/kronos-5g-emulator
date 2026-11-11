// Constants of package logger.

package logger

// Log levels, from the most to the least verbose.
const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

// Output modes.
const (
	OutputAuto OutputMode = iota
	OutputConsole
	OutputFile
	OutputBoth
)

// The log file.
const (
	// logDirPerm is the permission of the log directory.
	logDirPerm = 0o755
	// logFilePerm is the permission of the log file.
	logFilePerm = 0o644
)

// Caller lookup.
const (
	// logfCallerSkip is the number of stack frames from callerInfo to the
	// application caller: callerInfo -> logf -> exported logging function -> caller.
	logfCallerSkip = 3
)
