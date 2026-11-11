package logger

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// slogLevel maps a slog level to the logger's level and label.
func slogLevel(level slog.Level) (Level, string) {
	switch {
	case level < slog.LevelDebug:
		return LevelTrace, "TRACE"
	case level < slog.LevelInfo:
		return LevelDebug, "DEBUG"
	case level < slog.LevelWarn:
		return LevelInfo, "INFO"
	case level < slog.LevelError:
		return LevelWarn, "WARN"
	default:
		return LevelError, "ERROR"
	}
}

// Enabled reports whether level passes the current log level.
func (*slogHandler) Enabled(_ context.Context, level slog.Level) bool {
	v, _ := slogLevel(level)
	mu.Lock()
	defer mu.Unlock()
	return v >= currentLevel
}

// Handle writes the record in the common format, with its attributes appended as key=value.
func (h *slogHandler) Handle(_ context.Context, r slog.Record) error {
	level, label := slogLevel(r.Level)
	var message strings.Builder
	message.WriteString(r.Message)
	message.WriteString(h.attrs)
	r.Attrs(func(a slog.Attr) bool { appendAttr(&message, h.group, a); return true })
	function, file, line := "unknown", "unknown", 0
	if r.PC != 0 {
		frame, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()
		function, file, line = shortFunctionName(frame.Function), filepath.Base(frame.File), frame.Line
	}
	at := r.Time
	if at.IsZero() {
		at = time.Now()
	}
	mu.Lock()
	defer mu.Unlock()
	if level >= currentLevel {
		writeLocked(at, label, function, file, line, message.String())
	}
	return nil
}

// WithAttrs returns a copy with attrs formatted into its message suffix.
func (h *slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	next := *h
	var b strings.Builder
	b.WriteString(h.attrs)
	for _, a := range attrs {
		appendAttr(&b, h.group, a)
	}
	next.attrs = b.String()
	return &next
}

// WithGroup returns a copy that prefixes later attribute keys with name.
func (h *slogHandler) WithGroup(name string) slog.Handler {
	next := *h
	if name != "" {
		next.group += name + "."
	}
	return &next
}

// appendAttr writes a as " key=value", flattening groups into dotted keys and skipping empty attrs.
func appendAttr(b *strings.Builder, group string, a slog.Attr) {
	a.Value = a.Value.Resolve()
	if a.Equal(slog.Attr{}) {
		return
	}
	if a.Value.Kind() == slog.KindGroup {
		if a.Key != "" {
			group += a.Key + "."
		}
		for _, child := range a.Value.Group() {
			appendAttr(b, group, child)
		}
		return
	}
	b.WriteByte(' ')
	b.WriteString(attrText(group + a.Key))
	b.WriteByte('=')
	b.WriteString(attrText(fmt.Sprint(a.Value.Any())))
}

// attrText quotes value when it is empty or holds spaces, '=', quotes or backslashes.
func attrText(value string) string {
	if value == "" || strings.ContainsAny(value, " \t\r\n=\"\\") {
		return strconv.Quote(value)
	}
	return value
}
