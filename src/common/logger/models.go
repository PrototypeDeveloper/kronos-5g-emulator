// Structs of package logger.

package logger

// Compatibility for dependencies using slog or the standard log package.
// WithAttrs/WithGroup return immutable handlers safe to share across goroutines.
type slogHandler struct {
	attrs string
	group string
}
