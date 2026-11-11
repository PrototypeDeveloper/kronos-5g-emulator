// Package logger is the shared leveled logger. It writes
// "[time][LEVEL][func][file:line] message" lines to the console and/or
// log/<name>.txt, and can route slog and the standard log package there too.
package logger

import (
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Level is a log severity; messages below the current level are dropped.
type Level int

// OutputMode selects where log lines go.
type OutputMode int

var (
	mu           sync.Mutex
	currentLevel           = LevelDebug
	writer       io.Writer = os.Stdout
	logFile      *os.File
	serviceName  string
	outputMode   = OutputAuto
)

// Init initializes the common logger.
//
// Auto mode behavior:
//   - debug/trace level: console and ./log/<serviceName>.txt
//   - other levels, go test: console
//   - other levels, go run: console (detected from the temporary go-build executable path)
//   - other levels, built binary: ./log/<serviceName>.txt
//
// CORE_LOG_OUTPUT can override auto detection with: auto, console (or stdout), file, both.
// Set the level before Init; Configure does this at service startup.
func Init(name string) error {
	return InitWithMode(name, OutputAuto)
}

// InitWithMode is Init with an explicit output mode; CORE_LOG_OUTPUT still overrides it.
func InitWithMode(name string, mode OutputMode) error {
	mu.Lock()
	defer mu.Unlock()

	serviceName = name
	outputMode = resolveOutputMode(mode)

	if logFile != nil {
		_ = logFile.Close()
		logFile = nil
	}
	// Initialization failures must still be reportable after closing an old file.
	writer = os.Stdout

	switch outputMode {
	case OutputConsole:
		writer = os.Stdout
		return nil

	case OutputFile, OutputBoth:
		if strings.TrimSpace(serviceName) == "" {
			return fmt.Errorf("logger: service name is empty")
		}

		// Append to log/<name>.txt under the current directory.
		if err := os.MkdirAll("log", logDirPerm); err != nil {
			return fmt.Errorf("logger: create log directory: %w", err)
		}

		path := filepath.Join("log", serviceName+".txt")
		f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, logFilePerm)
		if err != nil {
			return fmt.Errorf("logger: open log file %s: %w", path, err)
		}

		logFile = f
		writer = f
		if outputMode == OutputBoth {
			writer = io.MultiWriter(f, os.Stdout)
		}
		return nil

	default:
		return fmt.Errorf("logger: invalid output mode: %d", outputMode)
	}
}

// Close closes the log file, if any, and sends output back to stdout.
func Close() error {
	mu.Lock()
	defer mu.Unlock()

	if logFile == nil {
		return nil
	}

	err := logFile.Close()
	logFile = nil
	writer = os.Stdout
	return err
}

// SetLevel sets the lowest level that is written.
func SetLevel(level Level) {
	mu.Lock()
	currentLevel = level
	mu.Unlock()
}

// SetLevelString sets the level from its name (see ParseLevel).
func SetLevelString(level string) error {
	parsed, err := ParseLevel(level)
	if err != nil {
		return err
	}
	SetLevel(parsed)
	return nil
}

// ParseLevel parses trace, debug, info, warn(ing) or error, ignoring case and spaces.
func ParseLevel(level string) (Level, error) {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "trace":
		return LevelTrace, nil
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn", "warning":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	default:
		return LevelDebug, fmt.Errorf("logger: unknown log level %q", level)
	}
}

// Error logs a printf-style message at error level.
func Error(format string, args ...any) {
	logf(LevelError, "ERROR", format, args...)
}

// Warn logs a printf-style message at warning level.
func Warn(format string, args ...any) {
	logf(LevelWarn, "WARN", format, args...)
}

// Debug logs a printf-style message at debug level.
func Debug(format string, args ...any) {
	logf(LevelDebug, "DEBUG", format, args...)
}

// Trace logs a printf-style message at trace level.
func Trace(format string, args ...any) {
	logf(LevelTrace, "TRACE", format, args...)
}

// logf writes one line tagged with the caller's position when level is enabled.
func logf(level Level, label string, format string, args ...any) {
	mu.Lock()
	defer mu.Unlock()

	if level < currentLevel {
		return
	}

	// callerInfo -> logf -> exported logging function -> application caller.
	funcName, fileName, line := callerInfo(logfCallerSkip)
	message := fmt.Sprintf(format, args...)
	writeLocked(time.Now(), label, funcName, fileName, line, message)
}

// Both native calls and compatibility handlers use the same writer and format.
func writeLocked(at time.Time, label, function, file string, line int, message string) {
	_, _ = fmt.Fprintf(
		writer,
		"[%s][%s][%s][%s:%d] %s\n",
		at.Format("2006-01-02 15:04:05.000"),
		label,
		function,
		file,
		line,
		message,
	)
}

// callerInfo returns the short function name, file base name and line skip frames up.
func callerInfo(skip int) (function string, file string, line int) {
	pc, filePath, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", "unknown", 0
	}

	fn := runtime.FuncForPC(pc)
	function = "unknown"
	if fn != nil {
		function = shortFunctionName(fn.Name())
	}

	return function, filepath.Base(filePath), line
}

func shortFunctionName(name string) string {
	// Keep package/function while dropping a long module path.
	// Example: kronos/simulator/engine.(*Engine).Reload -> engine.(*Engine).Reload
	if idx := strings.LastIndex(name, "/"); idx >= 0 {
		return name[idx+1:]
	}
	return name
}

// resolveOutputMode applies CORE_LOG_OUTPUT and resolves OutputAuto as described at Init.
func resolveOutputMode(requested OutputMode) OutputMode {
	if env := strings.ToLower(strings.TrimSpace(os.Getenv("CORE_LOG_OUTPUT"))); env != "" {
		switch env {
		case "auto":
			requested = OutputAuto
		case "console", "stdout":
			return OutputConsole
		case "file":
			return OutputFile
		case "both":
			return OutputBoth
		}
	}

	if requested != OutputAuto {
		return requested
	}

	if currentLevel <= LevelDebug {
		return OutputBoth
	}

	if isGoTest() || isGoRun() {
		return OutputConsole
	}

	return OutputFile
}

// isGoTest reports whether the process is a go test binary.
func isGoTest() bool {
	return flag.Lookup("test.v") != nil || strings.HasSuffix(os.Args[0], ".test")
}

// isGoRun reports whether the binary was built by go run (it sits in a go-build temp directory).
func isGoRun() bool {
	exe, err := os.Executable()
	if err != nil {
		return false
	}

	normalized := filepath.ToSlash(exe)
	return strings.Contains(normalized, "/go-build")
}

// Configure routes common, slog and standard log calls through one formatter.
// Call once at startup, before launching service goroutines.
func Configure(name, level string) error {
	parsed, err := ParseLevel(level)
	if err != nil {
		return err
	}
	SetLevel(parsed)
	if err := Init(name); err != nil {
		return err
	}
	// SetDefault captures this flag before disabling standard log's own prefix.
	log.SetFlags(log.Lshortfile)
	slog.SetDefault(slog.New(&slogHandler{}))
	return nil
}

// Info logs a printf-style message at info level.
func Info(format string, args ...any) {
	logf(LevelInfo, "INFO", format, args...)
}
