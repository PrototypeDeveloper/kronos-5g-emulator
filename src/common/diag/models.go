// Structs of package diag.

package diag

// Pos is a source position. Line and Col are 1-based; zero means unknown.
type Pos struct {
	File string
	Line int
	Col  int
}

// Diag is one diagnostic. It implements error so parsers can return it directly.
type Diag struct {
	Sev Severity
	Pos Pos
	Msg string
}
