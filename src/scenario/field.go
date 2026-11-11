package scenario

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"

	libnas "libnas"
	"libnas/nasMessage"

	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/command"
)

// Reading an IE of a received message for $nas.X / $ngap.X: each name of
// the path is looked for under the IE before it (its own IEs first, then
// deeper), only through what the message has (optional IEs that are
// present, the alternative a CHOICE took); [index] picks a list item.

// FieldValue returns the IE that path names in the first of roots (decoded
// NGAP or NAS messages) that has it; a message without it gives TypeNone.
func FieldValue(roots []any, path []command.FieldSeg) (command.Value, error) {
	for _, r := range roots {
		if r == nil {
			continue
		}
		v, parent, name, ok := fieldWalk(reflect.ValueOf(r), path)
		if ok {
			return fieldLeaf(v, parent, name, path)
		}
	}
	return command.Value{}, nil
}

// NASMessages decodes plain NAS messages for FieldValue: the 5GSM message a
// UL/DL NAS Transport carries comes before the transport itself.
func NASMessages(plains [][]byte) []any {
	var out []any
	for _, p := range plains {
		m, err := libnas.Decode(p, nil)
		if err != nil {
			continue
		}
		var c []byte
		switch t := m.(type) {
		case *nasMessage.DLNASTransport:
			if t.PayloadContainerType == payloadN1SMInfo {
				c = t.PayloadContainer
			}
		case *nasMessage.ULNASTransport:
			if t.PayloadContainerType == payloadN1SMInfo {
				c = t.PayloadContainer
			}
		}
		if c != nil {
			if inner, err := libnas.Decode(c, nil); err == nil {
				out = append(out, inner)
			}
		}
		out = append(out, m)
	}
	return out
}

// fieldWalk follows path from v; it returns the IE, the struct that holds it and its field name.
func fieldWalk(v reflect.Value, path []command.FieldSeg) (reflect.Value, reflect.Type, string, bool) {
	var parent reflect.Type
	var name string
	for i, seg := range path {
		if i > 0 && !path[i-1].HasIndex {
			v = openTransfer(v, name) // a name after a transfer is looked for inside it
		}
		f, pt, n, ok := fieldFind(v, seg.Name)
		if !ok {
			return v, nil, "", false
		}
		if seg.HasIndex {
			if f, ok = fieldItem(f, seg.Index); !ok {
				return v, nil, "", false
			}
		}
		v, parent, name = f, pt, n
	}
	return v, parent, name, true
}

// openTransfer decodes a transfer (an NGAP value carried in octets, such as
// a PDUSessionResourceSetupRequestTransfer) so that its IEs can be named;
// other values are returned as they are.
func openTransfer(v reflect.Value, name string) reflect.Value {
	typ, ok := transferTypes[name]
	if !ok || v.Kind() != reflect.Slice || v.Len() == 0 {
		return v
	}
	p := reflect.New(typ)
	if err := ngapcodec.DecodeParameter(typ.Name(), v.Bytes(), p.Interface()); ngapcodec.IsFatal(err) {
		return v
	}
	return p.Elem()
}

// presentFields lists the IEs of a struct value that the message has, in
// order: of a CHOICE or an IE entry only the alternative it took.
func presentFields(v reflect.Value) []int {
	t := v.Type()
	if alts, at := choiceOf(t); alts != nil {
		if f, ok := t.FieldByName(chosen(v, at, alts)); ok {
			return f.Index[:1]
		}
		return nil
	}
	var out []int
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() || skippedField[sf.Name] || isHasFlag(t, sf) {
			continue
		}
		if has := v.FieldByName("Has" + sf.Name); has.IsValid() && has.Kind() == reflect.Bool && !has.Bool() {
			continue
		}
		out = append(out, i)
	}
	return out
}

// skippedField are the struct fields that are no IEs of their own.
var skippedField = map[string]bool{"UnknownExtensions": true, "AdditionalIEs": true, "UnknownIEs": true, "UnknownValue": true}

// find looks for the IE called name under v: its own IEs first, then deeper.
func fieldFind(v reflect.Value, name string) (reflect.Value, reflect.Type, string, bool) {
	v = fieldIndirect(v)
	switch v.Kind() {
	case reflect.Struct:
		idx := presentFields(v)
		for _, i := range idx {
			if sf := v.Type().Field(i); strings.EqualFold(sf.Name, name) {
				return v.Field(i), v.Type(), sf.Name, true
			}
		}
		for _, i := range idx {
			if f, pt, n, ok := fieldFind(openTransfer(v.Field(i), v.Type().Field(i).Name), name); ok {
				return f, pt, n, true
			}
		}
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			break
		}
		for i := 0; i < v.Len(); i++ {
			if f, pt, n, ok := fieldFind(v.Index(i), name); ok {
				return f, pt, n, true
			}
		}
	}
	return v, nil, "", false
}

// item returns the item at i of a list IE (a slice, or a struct holding one in List).
func fieldItem(v reflect.Value, i int) (reflect.Value, bool) {
	v = fieldIndirect(v)
	if v.Kind() == reflect.Struct {
		v = v.FieldByName("List")
	}
	if !v.IsValid() || v.Kind() != reflect.Slice || v.Type().Elem().Kind() == reflect.Uint8 || i >= v.Len() {
		return v, false
	}
	return v.Index(i), true
}

// indirect follows pointers.
func fieldIndirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return reflect.Value{}
		}
		v = v.Elem()
	}
	return v
}

// fieldLeaf turns an IE into a value: numbers stay numbers, NGAP ENUMERATED
// values become their names, octets are hex, a CHOICE is the name of the
// alternative it took.
func fieldLeaf(v reflect.Value, parent reflect.Type, name string, path []command.FieldSeg) (command.Value, error) {
	v = fieldIndirect(v)
	what := strings.TrimPrefix(command.Expr{Kind: command.ExprField, Path: path}.String(), "$.")
	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		if alts, at := choiceOf(t); alts != nil {
			if alt := chosen(v, at, alts); alt != "" {
				return command.StrVal(alt), nil
			}
			return command.Value{}, nil
		}
		if t == bitStringType {
			return bitsValue(v.Interface().(ngapType.BitString), parent), nil
		}
		if idx := presentFields(v); len(idx) == 1 {
			return fieldLeaf(v.Field(idx[0]), t, t.Field(idx[0]).Name, path)
		}
		return command.Value{}, fmt.Errorf("%s is a group of IEs: name one of them", what)
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return command.StrVal(hex.EncodeToString(v.Bytes())), nil
		}
		return command.Value{}, fmt.Errorf("%s is a list: pick an item with [index]", what)
	case reflect.Array:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			b := make([]byte, v.Len())
			reflect.Copy(reflect.ValueOf(b), v)
			return command.StrVal(hex.EncodeToString(b)), nil
		}
	case reflect.String:
		return command.StrVal(v.String()), nil
	case reflect.Bool:
		return command.BoolVal(v.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return command.IntVal(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n := v.Uint()
		for _, o := range enumOptions(parent, name) {
			if uint64(o.Value) == n {
				return command.StrVal(o.Name), nil
			}
		}
		return command.IntVal(int64(n)), nil
	}
	return command.Value{}, fmt.Errorf("%s cannot be read", what)
}

// bitsValue reads a BIT STRING as the signal editor shows it: the
// addresses of a TransportLayerAddress, a short one as a number, else hex.
func bitsValue(bs ngapType.BitString, parent reflect.Type) command.Value {
	if parent == transportType {
		return command.StrVal(transportText(bs))
	}
	if bs.BitLength < maxBitsAsNumber && uint64(len(bs.Bytes)*8) >= bs.BitLength { // an int64 holds 63 bits
		var x int64
		for i := uint64(0); i < bs.BitLength; i++ {
			x = x<<1 | int64(bs.Bytes[i/8]>>(7-i%8)&1)
		}
		return command.IntVal(x)
	}
	return command.StrVal(hex.EncodeToString(bs.Bytes))
}
