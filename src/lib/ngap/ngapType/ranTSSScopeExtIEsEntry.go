// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RANTSSScopeExtIEsEntryPresent uint32

const (
	RANTSSScopeExtIEsEntryPresentNothing RANTSSScopeExtIEsEntryPresent = iota
	RANTSSScopeExtIEsEntryPresentUnknown
)

type RANTSSScopeExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      RANTSSScopeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANTSSScopeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RAN-TSSScope-ExtIEsEntry", v)
}

// DecodeRANTSSScopeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRANTSSScopeExtIEsEntry(data []byte) (RANTSSScopeExtIEsEntry, int, error) {
	var v RANTSSScopeExtIEsEntry
	err := ngapcodec.DecodeParameter("RAN-TSSScope-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTSSScopeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
