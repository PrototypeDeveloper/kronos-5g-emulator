package scenario

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	libnas "libnas"
	"libnas/nasMessage"
	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/command"
	"kronos/common/diag"
	"kronos/models"
)

// Mode is how a signal is sent.
type Mode uint8

// String returns the mode as written in the signal options.
func (m Mode) String() string {
	switch m {
	case ModeNASKeep:
		return "nas=keep"
	case ModeRaw:
		return "raw"
	default:
		return "auto"
	}
}

// Label renders the part like "ULNASTransport>PDUSessionEstablishmentRequest(2)".
func (p NASPart) Label() string {
	s := p.Name
	if p.Inner != "" {
		s += ">" + p.Inner
	}
	if p.Header != 0 {
		s += fmt.Sprintf("(%d)", p.Header)
	}
	return s
}

// Label renders "NGAP[:NAS]" in the form receive() triggers use.
func (s *Signal) Label() string {
	if s.NGAP == nil {
		return "?"
	}
	if len(s.NAS) == 0 || s.NAS[0].Name == "" {
		return s.NGAP.Name
	}
	return s.NGAP.Name + ":" + s.NAS[0].Label()
}

// Lookup returns the signal named name (case-sensitive, as names are labels).
func (s *SignalSet) Lookup(name string) (*Signal, bool) {
	if s == nil {
		return nil, false
	}
	sig, ok := s.byName[name]
	return sig, ok
}

// Names returns the signal names in file order.
func (s *SignalSet) Names() []string {
	if s == nil {
		return nil
	}
	out := make([]string, len(s.List))
	for i, sig := range s.List {
		out[i] = sig.Name
	}
	return out
}

// ReadSignals reads and parses a signal file.
func ReadSignals(path string) (*SignalSet, diag.List) {
	b, err := os.ReadFile(path)
	if err != nil {
		var ds diag.List
		ds.Errorf(diag.Pos{File: path}, "cannot read the signal file: %v", err)
		return nil, ds
	}
	return ParseSignals(path, b)
}

// ParseSignals parses signal.txt and decodes every signal.
//
//	signal  := NAME [ "[" option { "," option } "]" ] "{" { hex | space | "#" comment } "}"
//	option  := "raw" | "nas=keep" | "nas=auto"
func ParseSignals(file string, src []byte) (*SignalSet, diag.List) {
	set := &SignalSet{Path: file, byName: map[string]*Signal{}}
	var ds diag.List
	sc := &scanner{src: src, file: file, line: 1, col: 1}
	for {
		sc.skipSpace()
		if sc.eof() {
			break
		}
		pos, start := sc.pos(), sc.off
		name := sc.word()
		if name == "" {
			ds.Errorf(pos, "a signal name is needed (found %q)", sc.peek())
			sc.skipPast('}')
			continue
		}
		sig := &Signal{Name: name, Pos: pos, Start: start}
		sc.skipSpace()
		// Optional [options] between the name and the body.
		if sc.peek() == '[' {
			sc.next()
			opts := sc.until(']')
			if !sc.eof() {
				sc.next()
			} else {
				ds.Errorf(pos, "%s: the ] of the options is missing", name)
			}
			sig.Mode = parseOptions(name, pos, opts, &ds)
			sc.skipSpace()
		}
		if sc.peek() != '{' {
			ds.Errorf(sc.pos(), "the signal name %s must be followed by {", name)
			sc.skipPast('}')
			continue
		}
		sc.next()
		// The hex body up to the closing }.
		body, ok := sc.hexBody(name, &ds)
		if !ok {
			continue
		}
		sig.Bytes, sig.End = body, sc.off
		// Register the signal under a unique name and decode it.
		if prev, dup := set.byName[name]; dup {
			ds.Errorf(pos, "the signal name %s is used twice (the first is on line %d)", name, prev.Pos.Line)
			continue
		}
		set.byName[name] = sig
		set.List = append(set.List, sig)
		analyze(sig, &ds)
	}
	return set, ds
}

// parseOptions reads the [options] of a signal into its send mode.
func parseOptions(name string, pos diag.Pos, opts string, ds *diag.List) Mode {
	mode, seen := ModeAuto, map[string]bool{}
	for _, o := range strings.Split(opts, ",") {
		o = strings.ToLower(strings.Join(strings.Fields(o), ""))
		switch o {
		case "raw":
			mode = ModeRaw
		case "nas=keep":
			mode = ModeNASKeep
		case "nas=auto":
		default:
			ds.Errorf(pos, "%s: unknown option %q (raw and nas=keep are allowed)", name, o)
			continue
		}
		seen[o] = true
	}
	if seen["raw"] && (seen["nas=keep"] || seen["nas=auto"]) {
		ds.Errorf(pos, "%s: raw and nas=… cannot go together (raw fills in nothing)", name)
	}
	return mode
}

func (s *scanner) eof() bool     { return s.off >= len(s.src) }
func (s *scanner) pos() diag.Pos { return diag.Pos{File: s.file, Line: s.line, Col: s.col} }

// peek returns the next byte without consuming it (0 at the end).
func (s *scanner) peek() byte {
	if s.eof() {
		return 0
	}
	return s.src[s.off]
}

// next consumes one byte, updating the line and column.
func (s *scanner) next() byte {
	c := s.src[s.off]
	s.off++
	if c == '\n' {
		s.line++
		s.col = 1
	} else if c&utf8ContMask != utf8ContBits { // count runes, not UTF-8 continuation bytes
		s.col++
	}
	return c
}

// skipSpace skips whitespace and '#' comments.
func (s *scanner) skipSpace() {
	for !s.eof() {
		switch c := s.peek(); {
		case c == ' ' || c == '\t' || c == '\r' || c == '\n':
			s.next()
		case c == '#':
			for !s.eof() && s.peek() != '\n' {
				s.next()
			}
		case c == utf8BOM[0] && bytes.HasPrefix(s.src[s.off:], []byte(utf8BOM)):
			s.next()
			s.next()
			s.next()
		default:
			return
		}
	}
}

// word reads a name made of letters, digits and '_'.
func (s *scanner) word() string {
	start := s.off
	for !s.eof() {
		c := s.peek()
		if c == '_' || (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			s.next()
			continue
		}
		break
	}
	return string(s.src[start:s.off])
}

// until reads up to end or the end of the line, leaving that byte unread.
func (s *scanner) until(end byte) string {
	start := s.off
	for !s.eof() && s.peek() != end && s.peek() != '\n' {
		s.next()
	}
	return string(s.src[start:s.off])
}

// skipPast skips up to and including end.
func (s *scanner) skipPast(end byte) {
	for !s.eof() {
		if s.next() == end {
			return
		}
	}
}

// hexVal returns the value of a hex digit.
func hexVal(c byte) (byte, bool) {
	switch {
	case c >= '0' && c <= '9':
		return c - '0', true
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10, true
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10, true
	}
	return 0, false
}

// hexBody reads hex digits up to the closing '}'.
func (s *scanner) hexBody(name string, ds *diag.List) ([]byte, bool) {
	var out []byte
	nibbles, ok := 0, true
	start := s.pos()
	for {
		s.skipSpace()
		if s.eof() {
			ds.Errorf(start, "%s is not closed (} is missing)", name)
			return nil, false
		}
		c := s.peek()
		if c == '}' {
			s.next()
			break
		}
		v, isHex := hexVal(c)
		if !isHex {
			ds.Errorf(s.pos(), "%s: %q is not a hex digit (comments start with #)", name, rune(c))
			ok = false
			s.next()
			continue
		}
		// Two digits make an octet, the first one high.
		if nibbles%2 == 0 {
			out = append(out, v<<4)
		} else {
			out[len(out)-1] |= v
		}
		nibbles++
		s.next()
	}
	// Reject a body with bad digits, no digits or an odd number of them.
	switch {
	case !ok:
		return nil, false
	case nibbles == 0:
		ds.Errorf(start, "%s is empty", name)
		return nil, false
	case nibbles%2 != 0:
		ds.Errorf(start, "%s: an odd number of hex digits (%d)", name, nibbles)
		return nil, false
	}
	return out, true
}

// analyze decodes a signal and its NAS-PDUs and records problems.
func analyze(sig *Signal, ds *diag.List) {
	sig.Scope = command.ScopeNode | command.ScopeUE
	fatal := diag.Error
	if sig.Mode != ModeAuto {
		fatal = diag.Warning
	}
	msg, err := libngap.Decode(sig.Bytes)
	if msg == nil {
		ds.Add(fatal, sig.Pos, "%s: cannot decode as NGAP: %v", sig.Name, err)
		return
	}
	if err != nil {
		ds.Warnf(sig.Pos, "%s: decoded with a notice: %v", sig.Name, err)
	}
	sig.Msg = msg
	sig.NGAP, _ = NGAPOf(msg)
	sig.Scope = ScopeOf(msg)
	if re, err := libngap.Encode(msg); err != nil || !bytes.Equal(re, sig.Bytes) {
		ds.Warnf(sig.Pos, "%s: encoding again does not give the same octets (fields that are not filled in may change too)", sig.Name)
	}
	for _, f := range findNAS(msg) {
		part, err := analyzeNAS(f.b)
		part.Path = f.path
		sig.NAS = append(sig.NAS, part)
		switch {
		case err == nil:
		case errors.Is(err, errUnsupportedNAS):
			if sig.Mode == ModeAuto {
				ds.Warnf(sig.Pos, "%s: libnas does not support %s, so its NAS is not filled in (it is sent as in the template)", sig.Name, part.Name)
			}
		default:
			ds.Add(fatal, sig.Pos, "%s: the NAS of %s: %v", sig.Name, f.path, err)
		}
	}
}

// ScopeOf returns ScopeUE when a decoded message carries UE NGAP IDs and
// ScopeNode otherwise. It settles the optional-ID messages (AssocEither)
// per message, both for templates and for received PDUs.
func ScopeOf(msg any) command.Scope {
	v := reflect.ValueOf(msg).Elem()
	for _, name := range ueIDFields {
		if !v.FieldByName(name).IsValid() {
			continue
		}
		if has := v.FieldByName("Has" + name); has.IsValid() && !has.Bool() {
			continue
		}
		return command.ScopeUE
	}
	return command.ScopeNode
}

// findNAS returns the non-empty NAS-PDUs of a decoded NGAP message.
func findNAS(msg any) []nasField {
	var out []nasField
	walk(reflect.ValueOf(msg), "", func(path string, v reflect.Value) bool {
		if v.Type() != nasPDUType {
			return true
		}
		if b := v.Field(0).Bytes(); len(b) > 0 {
			out = append(out, nasField{path: strings.TrimPrefix(path, "."), b: b})
		}
		return false
	})
	return out
}

// walk visits every value reachable from v, skipping optional IEs whose Has*
// flag is false. fn returns false to stop descending into a value.
func walk(v reflect.Value, path string, fn func(path string, v reflect.Value) bool) {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if !v.IsNil() {
			walk(v.Elem(), path, fn)
		}
		return
	}
	if !fn(path, v) {
		return
	}
	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			f := t.Field(i)
			if !f.IsExported() {
				continue
			}
			if has := v.FieldByName("Has" + f.Name); has.IsValid() && has.Kind() == reflect.Bool && !has.Bool() {
				continue
			}
			walk(v.Field(i), path+"."+f.Name, fn)
		}
	case reflect.Slice, reflect.Array:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return
		}
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i), fmt.Sprintf("%s[%d]", path, i), fn)
		}
	}
}

var errUnsupportedNAS = errors.New("unsupported NAS message")

// analyzeNAS names a NAS-PDU and checks that its plaintext can be read: the
// builder decodes templates, so a protected NAS must have been captured with
// NEA0 (or be integrity protected only).
func analyzeNAS(b []byte) (NASPart, error) {
	var p NASPart
	plain := b
	// Skip the security header of a protected 5GMM message.
	if len(b) >= 2 && b[0] == models.EPD5GMM && b[1]&secHeaderTypeMask != 0 {
		p.Header = b[1] & secHeaderTypeMask
		if len(b) < protectedNASMinLen {
			return p, fmt.Errorf("protected NAS too short (%d octets)", len(b))
		}
		plain = b[secHeaderLen:]
	}
	ciphered := p.Header == shtIntegrityCiphered || p.Header == shtIntegrityCipheredNew
	name, err := nasName(plain)
	if err != nil {
		if ciphered {
			return p, fmt.Errorf("cannot read the body: a capture ciphered with other than NEA0 cannot be a template (%v)", err)
		}
		return p, err
	}
	p.Name = name.Name
	if !name.Supported {
		return p, errUnsupportedNAS
	}
	m, err := libnas.Decode(plain, nil)
	if err != nil {
		return p, fmt.Errorf("cannot decode %s: %v", p.Name, err)
	}
	// UL/DL NAS Transport may carry a 5GSM message: name and decode it too.
	var container []byte
	switch t := m.(type) {
	case *nasMessage.ULNASTransport:
		if t.PayloadContainerType == payloadN1SMInfo { // N1 SM information
			container = t.PayloadContainer
		}
	case *nasMessage.DLNASTransport:
		if t.PayloadContainerType == payloadN1SMInfo {
			container = t.PayloadContainer
		}
	}
	if container != nil {
		inner, err := nasName(container)
		if err != nil {
			return p, fmt.Errorf("the 5GSM of %s: %v", p.Name, err)
		}
		p.Inner = inner.Name
		if inner.Supported {
			if _, err := libnas.Decode(container, nil); err != nil {
				return p, fmt.Errorf("cannot decode the %[2]s of %[1]s: %[3]v", p.Name, p.Inner, err)
			}
		}
	}
	return p, nil
}

// nasName looks up a plain NAS message by its EPD and message type.
func nasName(b []byte) (*models.NASInfo, error) {
	if len(b) < gmmHeaderLen {
		return nil, fmt.Errorf("NAS too short (%d octets)", len(b))
	}
	var typ uint8
	switch b[0] {
	case models.EPD5GMM:
		switch {
		case b[1]&secHeaderTypeMask != 0:
			return nil, fmt.Errorf("plain 5GMM expected, but the security header type is %d", b[1]&secHeaderTypeMask)
		case b[1] != 0:
			return nil, fmt.Errorf("the spare bits of the security header octet are set (%#02x)", b[1])
		}
		typ = b[2]
	case models.EPD5GSM:
		if len(b) < gsmHeaderLen {
			return nil, fmt.Errorf("5GSM too short (%d octets)", len(b))
		}
		typ = b[3]
	default:
		return nil, fmt.Errorf("EPD 0x%02x is neither 5GMM nor 5GSM", b[0])
	}
	info, ok := NASByType(b[0], typ)
	if !ok {
		return nil, fmt.Errorf("unknown NAS message type 0x%02x", typ)
	}
	return info, nil
}

// unusedSignals lists the signals that the scenario never sends.
func unusedSignals(set *SignalSet, used map[string]bool) []string {
	var out []string
	for _, s := range set.List {
		if !used[s.Name] {
			out = append(out, s.Name)
		}
	}
	sort.Strings(out)
	return out
}

// NASPDUs returns the NAS-PDUs of a decoded NGAP message: the top-level one
// first, then those in lists, in order.
func NASPDUs(msg any) [][]byte {
	var out [][]byte
	for _, f := range findNAS(msg) {
		out = append(out, f.b)
	}
	return out
}

// UEIDs returns the UE NGAP IDs a decoded message carries, from the
// RAN-UE-NGAP-ID / AMF-UE-NGAP-ID IEs or the UE-NGAP-IDs choice.
func UEIDs(msg any) (ran int64, hasRAN bool, amf int64, hasAMF bool) {
	v := reflect.ValueOf(msg)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	present := func(name string) bool {
		has := v.FieldByName("Has" + name)
		return !has.IsValid() || has.Bool()
	}
	if f := v.FieldByName("RANUENGAPID"); f.IsValid() && present("RANUENGAPID") {
		ran, hasRAN = f.Interface().(ngapType.RANUENGAPID).Value, true
	}
	if f := v.FieldByName("AMFUENGAPID"); f.IsValid() && present("AMFUENGAPID") {
		amf, hasAMF = f.Interface().(ngapType.AMFUENGAPID).Value, true
	}
	if f := v.FieldByName("UENGAPIDs"); f.IsValid() && present("UENGAPIDs") {
		ids := f.Interface().(ngapType.UENGAPIDs)
		switch ids.Present {
		case ngapType.UENGAPIDsPresentUENGAPIDPair:
			ran, hasRAN = ids.UENGAPIDPair.RANUENGAPID.Value, true
			amf, hasAMF = ids.UENGAPIDPair.AMFUENGAPID.Value, true
		case ngapType.UENGAPIDsPresentAMFUENGAPID:
			amf, hasAMF = ids.AMFUENGAPID.Value, true
		}
	}
	return
}

// PagingKey identifies a UE by its 5G-S-TMSI (AMF set ID, AMF pointer and
// 5G-TMSI), as Paging names it and as the 5G-GUTI carries it.
func PagingKey(setID, pointer uint64, tmsi []byte) string {
	return fmt.Sprintf("%03x-%02x-%x", setID, pointer, tmsi)
}

// PagingKeyOf returns the key of the UE a Paging message is for.
func PagingKeyOf(msg any) (string, bool) {
	m, ok := msg.(*ngapMessage.Paging)
	if !ok || m.UEPagingIdentity.Present != ngapType.UEPagingIdentityPresentFiveGSTMSI {
		return "", false
	}
	s := m.UEPagingIdentity.FiveGSTMSI
	return PagingKey(bitStringValue(s.AMFSetID.Value), bitStringValue(s.AMFPointer.Value), s.FiveGTMSI.Value), true
}
