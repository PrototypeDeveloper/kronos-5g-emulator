// Struct types of package scenario.

package scenario

import (
	"reflect"

	"kronos/command"
	"kronos/common/diag"
	"kronos/models"
)

// The parsed scenario.txt.

// File is a parsed scenario.txt.
type File struct {
	Path   string
	Vars   []*Var
	Node   *Block
	Any    *Block
	States []*Block
	states map[string]*Block
}

// Var is a declaration in the var block.
type Var struct {
	Pos    diag.Pos
	Public bool // shared by the whole pseudo-call; private vars are per subscriber
	Name   string
	Init   command.Value
}

// Block is a node, any or state block.
type Block struct {
	Kind     BlockKind
	Name     string
	Pos      diag.Pos
	Entry    *Entry
	Handlers []*Handler
}

// Entry runs when the node connects (node) or a subscriber enters the state.
type Entry struct {
	Pos   diag.Pos
	Calls []command.Call
}

// Handler is receive(trigger) { calls }.
type Handler struct {
	Pos     diag.Pos
	Trigger Trigger
	Calls   []command.Call
}

// Trigger is `NGAP[:NAS]` as written in receive().
type Trigger struct {
	Pos      diag.Pos
	NGAP     string
	NAS      string
	Wildcard bool
	// Resolved by Validate.
	NGAPInfo *NGAPInfo
	NASInfo  *models.NASInfo
}

// parser builds a File from the tokens of scenario.txt, collecting diagnostics.
type parser struct {
	lx *command.Lexer
	f  *File
	ds diag.List
}

// Validation of a scenario.

// validator checks a scenario and resolves the names its commands use.
type validator struct {
	f     *File
	sigs  *SignalSet
	vars  map[string]*Var
	ds    diag.List
	sent  map[string]bool // signals sent by the scenario
	moved map[string]bool // states that some move() targets
	noted map[string]bool // spelling notes already given
}

// The parsed signal.txt and its NAS-PDUs.

// Signal is one named NGAP template in signal.txt.
type Signal struct {
	Name  string
	Pos   diag.Pos
	Mode  Mode
	Bytes []byte
	Msg   any       // decoded NGAP message (pointer); nil if it could not be decoded
	NGAP  *NGAPInfo // nil if Msg is nil
	// Scope is ScopeUE when the template carries UE NGAP IDs and ScopeNode
	// otherwise; both bits when the signal could not be decoded.
	Scope command.Scope
	NAS   []NASPart
	// Start and End are the octet offsets of the signal in its file, from
	// its name to its closing brace (included).
	Start, End int
}

// NASPart is one NAS-PDU found inside a signal.
type NASPart struct {
	Path   string // field path inside the NGAP message
	Header uint8  // 5GMM security header type (0 = plain)
	Name   string // NAS message name, "" if unknown
	Inner  string // the 5GSM message carried by UL/DL NAS Transport, if any
}

// SignalSet is the parsed signal.txt.
type SignalSet struct {
	Path   string
	List   []*Signal
	byName map[string]*Signal
}

// scanner reads signal.txt character by character, tracking line and column.
type scanner struct {
	src       []byte
	file      string
	off       int
	line, col int
}

// nasField is a NAS-PDU found in an NGAP message, with its field path.
type nasField struct {
	path string
	b    []byte
}

// NGAP message table.

// NGAPInfo describes one NGAP message type.
type NGAPInfo struct {
	Name       string       // libngap type name, e.g. "DownlinkNASTransport"
	Type       reflect.Type // the struct type
	Assoc      Assoc
	CarriesNAS bool // a NAS-PDU can appear somewhere inside
}

// NAS message table.

// Comparison of templates with the config.

// NGAP tree for the web console.

// treeBuilder renders decoded NGAP and NAS values as console tree nodes.
type treeBuilder struct {
	plain    map[string][]byte // NAS-PDU (as sent) → plaintext
	template bool
}

// The signal editor.

// FieldRef names a field for the auto-fix rules of a role.
type FieldRef struct {
	Mode  Mode
	NGAP  string // the NGAP message
	NAS   string // the NAS message the field is in ("" for NGAP fields)
	Outer string // the NAS message whose container holds that NAS message ("" if none)
	Path  string // relative to the innermost message ("@MAC" for a NAS security header field)
}

// loc is a place in a decoded signal.
type loc struct {
	v      reflect.Value       // addressable
	parent reflect.Value       // the struct or the list that holds v (invalid at a root)
	field  reflect.StructField // the field of v, or of its list for a list item
	item   bool                // v is an item of the list parent
	index  int
}

// nasHeader is the security header of a protected NAS-PDU, edited as fields.
type nasHeader struct {
	SecurityHeaderType uint8  `aper:"valueLB:0,valueUB:15"`
	MAC                []byte `aper:"sizeLB:4,sizeUB:4"`
	SequenceNumber     uint8
}

// container is the message held by octets, opened for editing.
type container struct {
	kind   string // nas-pdu, nas, transfer
	root   loc
	name   string     // the message's type name
	header *nasHeader // a protected NAS-PDU's header
	close  func() ([]byte, error)
}

// aperTag holds the constraints of an aper struct tag.
type aperTag struct {
	lb, ub, sizeLB, sizeUB *int64
	ext, sizeExt           bool
}

// viewer builds the editor tree of one signal, asking auto which fields a role replaces.
type viewer struct {
	mode Mode
	auto AutoRule
	ngap string
}

// ctx is where a node is: in which NAS message, and from which path its
// relative path starts.
type ctx struct {
	nas, outer string
	base       string // the path of the innermost NAS message ("" for NGAP fields)
}
