// Constants of package diag.

package diag

// Severities, from the most to the least serious.
const (
	Error   Severity = iota // the file cannot be used as written
	Warning                 // usable, but probably not what the author meant
	Info                    // for the author's information only
)
