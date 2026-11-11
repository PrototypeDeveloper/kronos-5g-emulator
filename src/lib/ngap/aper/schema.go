package aper

// Type describes PER-visible ASN.1 structure independently of any Go library.
// A Codec owns its schema; callers must not mutate it during use.
type Type struct {
	Kind        string
	Ref         string
	Constraints []Constraint
	Fields      []Field
	Values      []EnumValue
	Element     *Type
	Extensible  bool
	RootCount   int
}
type Constraint struct {
	Kind, Lower, Upper, Containing string
	Extensible                     bool
	Ranges, ExtensionRanges        []ValueRange
}
type ValueRange struct{ Lower, Upper string }
type Field struct {
	Name, GoField string
	Type          *Type
	Optional      bool
	Default       string
	Extension     bool
}
type EnumValue struct {
	Name      string
	Number    int64
	Extension bool
}
type IE struct {
	ID                             int64
	GoField, Criticality, Presence string
	Type                           *Type
}
type ObjectSet struct {
	Class string
	IEs   []IE
}
type Schema struct {
	Types map[string]*Type
	Names map[string]string
	Sets  map[string]ObjectSet
}

// Limits bounds untrusted input, nested open types and reflection allocations.
type Limits struct{ MaxBytes, MaxElements, MaxDepth, MaxAllocation int }

func DefaultLimits() Limits { return Limits{16 << 20, 65536, 128, 64 << 20} }

// Diagnostic reports an unknown IE with notify criticality. The decoded value
// remains usable; the application decides how to send NGAP Error Indication.
type Diagnostic struct {
	ID          int64
	Criticality string
	Problem     string
}
type NotifyError struct{ Diagnostics []Diagnostic }

func (e *NotifyError) Error() string { return "NGAP decoded with notify-criticality diagnostics" }

// CriticalityError requires rejecting the decoded procedure. It never returns
// a partially decoded value through Codec.Unmarshal.
type CriticalityError struct {
	ID      int64
	Problem string
}

func (e *CriticalityError) Error() string { return "NGAP reject-criticality IE: " + e.Problem }
