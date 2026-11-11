// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CriticalityDiagnosticsIEItemExtIEsEntryPresent uint32

const (
	CriticalityDiagnosticsIEItemExtIEsEntryPresentNothing CriticalityDiagnosticsIEItemExtIEsEntryPresent = iota
	CriticalityDiagnosticsIEItemExtIEsEntryPresentUnknown
)

type CriticalityDiagnosticsIEItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CriticalityDiagnosticsIEItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CriticalityDiagnosticsIEItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CriticalityDiagnostics-IE-Item-ExtIEsEntry", v)
}

// DecodeCriticalityDiagnosticsIEItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCriticalityDiagnosticsIEItemExtIEsEntry(data []byte) (CriticalityDiagnosticsIEItemExtIEsEntry, int, error) {
	var v CriticalityDiagnosticsIEItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CriticalityDiagnostics-IE-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CriticalityDiagnosticsIEItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
