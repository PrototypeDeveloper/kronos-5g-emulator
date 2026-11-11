package scenario

import (
	"reflect"
	"strconv"
	"strings"
	"sync"

	"libngap/aper"
	"libngap/ngapcodec"

	"kronos/models"
)

// What libngap's ASN.1 schema tells the editor: the names of ENUMERATED
// values, and the IEs of IE containers (protocol IEs, extensions).

var (
	schemaOnce sync.Once
	ngapSchema aper.Schema
)

func schema() aper.Schema {
	schemaOnce.Do(func() { ngapSchema = ngapcodec.NewSchema() })
	return ngapSchema
}

// asnType resolves a schema type through its references.
func asnType(t *aper.Type) *aper.Type {
	for i := 0; t != nil && t.Kind == "REFERENCE" && i < maxSchemaRefs; i++ {
		t = schema().Types[t.Ref]
	}
	return t
}

// enumOptions returns the named values of an ENUMERATED field: field of the
// Go struct parent (Value of a wrapper such as RRCEstablishmentCause, or an
// enumerated field of a SEQUENCE or CHOICE); nil when it is not one.
func enumOptions(parent reflect.Type, field string) []models.EnumOption {
	if parent == nil || parent.Name() == "" || parent.PkgPath() != ngapTypePkg {
		return nil // NAS has types of the same names: only NGAP's are in the schema
	}
	s := schema()
	t := asnType(s.Types[s.Names[parent.Name()]])
	if t == nil {
		return nil
	}
	if t.Kind != "ENUMERATED" {
		var ft *aper.Type
		for _, f := range t.Fields {
			if f.GoField == field {
				ft = asnType(f.Type)
			}
		}
		if ft == nil || ft.Kind != "ENUMERATED" {
			return nil
		}
		t = ft
	} else if field != "Value" {
		return nil
	}
	out := make([]models.EnumOption, 0, len(t.Values))
	for _, v := range t.Values {
		out = append(out, models.EnumOption{Value: v.Number, Name: v.Name, Ext: v.Extension})
	}
	return out
}

func parseIETag(tag reflect.StructTag) (ieTag, bool) {
	s := tag.Get("ngap")
	if s == "" {
		return ieTag{}, false
	}
	var t ieTag
	ok := false
	for _, part := range strings.Split(s, ",") {
		k, v, _ := strings.Cut(part, ":")
		switch k {
		case "id":
			n, err := strconv.ParseInt(v, 10, 64)
			t.id, ok = n, err == nil
		case "criticality":
			t.criticality = v
		case "presence":
			t.mandatory = v == "mandatory"
		}
	}
	return t, ok
}

// entryAlternatives returns the IEs an IE container entry can hold (its
// alternatives with an ngap tag); nil when there are none (an entry of
// only unknown values).
func entryAlternatives(typ reflect.Type) []reflect.StructField {
	if !isIEEntry(typ) {
		return nil
	}
	var out []reflect.StructField
	for i := 0; i < typ.NumField(); i++ {
		if _, ok := parseIETag(typ.Field(i).Tag); ok {
			out = append(out, typ.Field(i))
		}
	}
	return out
}

// editableEntries reports a list of IE container entries the editor can
// build (its entry type has IEs libngap knows).
func editableEntries(t reflect.Type) bool {
	return t.Kind() == reflect.Slice && len(entryAlternatives(t.Elem())) > 0
}

// criticalities maps the names of Criticality to its values.
var criticalities = map[string]int64{"reject": criticalityReject, "ignore": criticalityIgnore, "notify": criticalityNotify}

// setEntryIDs writes the IE ID and the criticality of the chosen IE of an
// entry, as its ngap tag gives them.
func setEntryIDs(v reflect.Value, alt string) {
	f, ok := v.Type().FieldByName(alt)
	if !ok {
		return
	}
	t, ok := parseIETag(f.Tag)
	if !ok {
		return
	}
	if id := v.FieldByName("ID"); id.IsValid() && id.Kind() == reflect.Struct && id.NumField() == 1 {
		setInt(id.Field(0), t.id)
	}
	if c := v.FieldByName("Criticality"); c.IsValid() && c.Kind() == reflect.Struct && c.NumField() == 1 {
		setInt(c.Field(0), criticalities[t.criticality])
	}
}

// entryHint describes the IE chosen in an entry: "id 139, reject".
func entryHint(typ reflect.Type, alt string) string {
	f, ok := typ.FieldByName(alt)
	if !ok {
		return ""
	}
	t, ok := parseIETag(f.Tag)
	if !ok {
		return ""
	}
	return "id " + strconv.FormatInt(t.id, 10) + ", " + t.criticality
}
