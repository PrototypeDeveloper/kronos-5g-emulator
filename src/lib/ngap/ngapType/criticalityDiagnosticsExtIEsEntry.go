// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CriticalityDiagnosticsExtIEsEntryPresent uint32

const (
	CriticalityDiagnosticsExtIEsEntryPresentNothing CriticalityDiagnosticsExtIEsEntryPresent = iota
	CriticalityDiagnosticsExtIEsEntryPresentUnknown
)

type CriticalityDiagnosticsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CriticalityDiagnosticsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CriticalityDiagnosticsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CriticalityDiagnostics-ExtIEsEntry", v)
}

// DecodeCriticalityDiagnosticsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCriticalityDiagnosticsExtIEsEntry(data []byte) (CriticalityDiagnosticsExtIEsEntry, int, error) {
	var v CriticalityDiagnosticsExtIEsEntry
	err := ngapcodec.DecodeParameter("CriticalityDiagnostics-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CriticalityDiagnosticsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
