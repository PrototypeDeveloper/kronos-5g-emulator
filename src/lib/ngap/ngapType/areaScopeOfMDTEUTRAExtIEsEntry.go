// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaScopeOfMDTEUTRAExtIEsEntryPresent uint32

const (
	AreaScopeOfMDTEUTRAExtIEsEntryPresentNothing AreaScopeOfMDTEUTRAExtIEsEntryPresent = iota
	AreaScopeOfMDTEUTRAExtIEsEntryPresentUnknown
)

type AreaScopeOfMDTEUTRAExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      AreaScopeOfMDTEUTRAExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfMDTEUTRAExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfMDT-EUTRA-ExtIEsEntry", v)
}

// DecodeAreaScopeOfMDTEUTRAExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfMDTEUTRAExtIEsEntry(data []byte) (AreaScopeOfMDTEUTRAExtIEsEntry, int, error) {
	var v AreaScopeOfMDTEUTRAExtIEsEntry
	err := ngapcodec.DecodeParameter("AreaScopeOfMDT-EUTRA-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfMDTEUTRAExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
