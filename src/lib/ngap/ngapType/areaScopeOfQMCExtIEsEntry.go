// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaScopeOfQMCExtIEsEntryPresent uint32

const (
	AreaScopeOfQMCExtIEsEntryPresentNothing AreaScopeOfQMCExtIEsEntryPresent = iota
	AreaScopeOfQMCExtIEsEntryPresentUnknown
)

type AreaScopeOfQMCExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      AreaScopeOfQMCExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfQMCExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfQMC-ExtIEsEntry", v)
}

// DecodeAreaScopeOfQMCExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfQMCExtIEsEntry(data []byte) (AreaScopeOfQMCExtIEsEntry, int, error) {
	var v AreaScopeOfQMCExtIEsEntry
	err := ngapcodec.DecodeParameter("AreaScopeOfQMC-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfQMCExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
