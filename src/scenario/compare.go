package scenario

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"libngap/ngapType"

	"kronos/common/diag"
	"kronos/common/plmn"
	"kronos/models"
)

var (
	plmnType = reflect.TypeOf(ngapType.PLMNIdentity{})
	nciType  = reflect.TypeOf(ngapType.NRCellIdentity{})
	tacType  = reflect.TypeOf(ngapType.TAC{})
	gnbIDTyp = reflect.TypeOf(ngapType.GNBID{})
)

// bitStringValue returns the BIT STRING as an integer, dropping the padding bits.
func bitStringValue(bs ngapType.BitString) uint64 {
	var v uint64
	for _, b := range bs.Bytes {
		v = v<<8 | uint64(b)
	}
	return v >> (uint64(len(bs.Bytes))*8 - bs.BitLength)
}

// Compare reports template values that auto-fix will overwrite (or, for raw
// signals, that will be sent although they differ from config).
func (s *SignalSet) Compare(ref models.SignalReference) diag.List {
	var ds diag.List
	if s == nil {
		return ds
	}
	want, wantOK := plmn.Encode(ref.MCC, ref.MNC)
	for _, sig := range s.List {
		if sig.Msg == nil {
			continue
		}
		fix := "the config values replace them when sent"
		if sig.Mode == ModeRaw {
			fix = "raw: sent as written"
		}
		var (
			badPLMN   = map[string]int{}
			otherPLMN = map[string]int{}
			diffs     []string
		)
		// Collect invalid PLMNs and PLMN, NCI, TAC and gNB ID values that differ from config.
		walk(reflect.ValueOf(sig.Msg), "", func(_ string, v reflect.Value) bool {
			switch v.Type() {
			case plmnType:
				b := v.Field(0).Bytes()
				if len(b) != plmnIDLen {
					return false
				}
				if mcc, mnc, ok := plmn.Decode(b); !ok {
					badPLMN[fmt.Sprintf("%x", b)]++
				} else if wantOK && !bytes.Equal(b, want) {
					otherPLMN[mcc+"/"+mnc]++
				}
				return false
			case nciType:
				bs := v.Field(0).Interface().(ngapType.BitString)
				if bs.BitLength == nciBitLen && ref.HasNCI {
					if got := bitStringValue(bs); got != ref.NCI {
						diffs = append(diffs, fmt.Sprintf("NR cell ID %09x (config %09x)", got, ref.NCI))
					}
				}
				return false
			case tacType:
				b := v.Field(0).Bytes()
				if len(b) == tacLen && ref.HasTAC {
					if got := uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2]); got != ref.TAC {
						diffs = append(diffs, fmt.Sprintf("TAC %d (config %d)", got, ref.TAC))
					}
				}
				return false
			case gnbIDTyp:
				bs := v.FieldByName("GNBID").Interface().(ngapType.BitString)
				if bs.BitLength > 0 && ref.HasGNBID {
					got := bitStringValue(bs)
					if got != uint64(ref.GNBID) || int(bs.BitLength) != ref.GNBIDLen {
						diffs = append(diffs, fmt.Sprintf("gNB ID %x/%d bits (config %x/%d bits)", got, bs.BitLength, ref.GNBID, ref.GNBIDLen))
					}
				}
				return false
			}
			return true
		})
		// Report the findings, one warning per kind.
		for _, k := range sortedKeys(badPLMN) {
			ds.Warnf(sig.Pos, "%s: PLMN %s is not valid BCD (%d place(s)); %s", sig.Name, k, badPLMN[k], fix)
		}
		if len(otherPLMN) > 0 {
			var parts []string
			for _, k := range sortedKeys(otherPLMN) {
				parts = append(parts, fmt.Sprintf("%s (%d place(s))", k, otherPLMN[k]))
			}
			ds.Warnf(sig.Pos, "%s: PLMN differs from %s/%s of config: %s; %s", sig.Name, ref.MCC, ref.MNC, strings.Join(parts, ", "), fix)
		}
		if len(diffs) > 0 {
			ds.Warnf(sig.Pos, "%s: values differ from config: %s; %s", sig.Name, strings.Join(dedup(diffs), ", "), fix)
		}
		// Without a configured name, the template's RANNodeName is sent as is.
		if ref.Name == "" && sig.NGAP != nil && sig.NGAP.Name == "NGSetupRequest" {
			v := reflect.ValueOf(sig.Msg).Elem()
			if v.FieldByName("HasRANNodeName").Bool() {
				ds.Infof(sig.Pos, "%s: RANNodeName is sent as %q of the template (gnodeb.name of config can change it)",
					sig.Name, v.FieldByName("RANNodeName").Field(0).String())
			}
		}
	}
	return ds
}

// sortedKeys returns the keys of m in sorted order.
func sortedKeys(m map[string]int) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// dedup removes repeated strings, keeping the first occurrence order.
func dedup(in []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, s := range in {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}
