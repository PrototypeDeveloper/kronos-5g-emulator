// Package scenario parses and validates scenario.txt and signal.txt.
// It also owns the protocol name tables: NGAP message
// names (from libngap) and 5GS NAS message names (TS 24.501 §9.7).
package scenario

import (
	"reflect"
	"sort"
	"strings"

	"libngap/ngapType"
)

// Assoc says whether an NGAP message belongs to a UE.
type Assoc uint8

func (a Assoc) String() string {
	switch a {
	case AssocUE:
		return "UE"
	case AssocEither:
		return "UE/node"
	default:
		return "node"
	}
}

var (
	nasPDUType  = reflect.TypeOf(ngapType.NASPDU{})
	ngapList    []*NGAPInfo
	ngapByLower = map[string]*NGAPInfo{}
	ngapByType  = map[reflect.Type]*NGAPInfo{}
)

// init builds the NGAP message tables from ngapMessageTypes.
func init() {
	for _, p := range ngapMessageTypes {
		t := reflect.TypeOf(p).Elem()
		info := &NGAPInfo{Name: t.Name(), Type: t, Assoc: classify(t), CarriesNAS: carriesNAS(t, map[reflect.Type]bool{})}
		ngapList = append(ngapList, info)
		ngapByLower[strings.ToLower(info.Name)] = info
		ngapByType[t] = info
	}
	sort.Slice(ngapList, func(i, j int) bool { return ngapList[i].Name < ngapList[j].Name })
}

// ueIDFields are the top-level IEs whose presence makes a message UE-associated.
var ueIDFields = []string{"AMFUENGAPID", "RANUENGAPID", "UENGAPIDs"}

// classify decides UE association from the UE NGAP ID fields of message type t.
func classify(t reflect.Type) Assoc {
	found, allOptional := false, true
	for _, name := range ueIDFields {
		f, ok := t.FieldByName(name)
		if !ok {
			continue
		}
		found = true
		if !strings.Contains(f.Tag.Get("ngap"), "presence:optional") {
			allOptional = false
		}
	}
	switch {
	case !found:
		return AssocNode
	case allOptional:
		return AssocEither
	default:
		return AssocUE
	}
}

// carriesNAS reports whether a NAS-PDU can appear anywhere inside type t.
func carriesNAS(t reflect.Type, seen map[reflect.Type]bool) bool {
	switch t.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Array:
		return carriesNAS(t.Elem(), seen)
	case reflect.Struct:
		if t == nasPDUType {
			return true
		}
		if seen[t] {
			return false
		}
		seen[t] = true
		for i := 0; i < t.NumField(); i++ {
			if carriesNAS(t.Field(i).Type, seen) {
				return true
			}
		}
	}
	return false
}

// LookupNGAP finds an NGAP message by name, ignoring case
// (so the sample's "DownLinkNasTransport" resolves to DownlinkNASTransport).
func LookupNGAP(name string) (*NGAPInfo, bool) {
	info, ok := ngapByLower[strings.ToLower(name)]
	return info, ok
}

// NGAPOf returns the info for a decoded message (a pointer to an ngapMessage struct).
func NGAPOf(msg any) (*NGAPInfo, bool) {
	t := reflect.TypeOf(msg)
	if t == nil {
		return nil, false
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	info, ok := ngapByType[t]
	return info, ok
}

// NGAPMessages returns every NGAP message, sorted by name.
func NGAPMessages() []*NGAPInfo { return ngapList }

// NGAPNames returns every NGAP message name, sorted.
func NGAPNames() []string {
	out := make([]string, len(ngapList))
	for i, m := range ngapList {
		out[i] = m.Name
	}
	return out
}
