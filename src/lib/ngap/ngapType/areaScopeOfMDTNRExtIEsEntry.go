// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaScopeOfMDTNRExtIEsEntryPresent uint32

const (
	AreaScopeOfMDTNRExtIEsEntryPresentNothing AreaScopeOfMDTNRExtIEsEntryPresent = iota
	AreaScopeOfMDTNRExtIEsEntryPresentPNINPNBasedMDT
	AreaScopeOfMDTNRExtIEsEntryPresentSNPNCellBasedMDT
	AreaScopeOfMDTNRExtIEsEntryPresentSNPNTAIBasedMDT
	AreaScopeOfMDTNRExtIEsEntryPresentSNPNBasedMDT
	AreaScopeOfMDTNRExtIEsEntryPresentGeographyBasedMDT
	AreaScopeOfMDTNRExtIEsEntryPresentUnknown
)

type AreaScopeOfMDTNRExtIEsEntry struct {
	ID                ProtocolIEID
	Criticality       Criticality
	Present           AreaScopeOfMDTNRExtIEsEntryPresent
	PNINPNBasedMDT    PNINPNBasedMDT    `ngap:"id:410,criticality:ignore,presence:mandatory"`
	SNPNCellBasedMDT  SNPNCellBasedMDT  `ngap:"id:411,criticality:ignore,presence:mandatory"`
	SNPNTAIBasedMDT   SNPNTAIBasedMDT   `ngap:"id:412,criticality:ignore,presence:mandatory"`
	SNPNBasedMDT      SNPNBasedMDT      `ngap:"id:413,criticality:ignore,presence:mandatory"`
	GeographyBasedMDT GeographyBasedMDT `ngap:"id:448,criticality:ignore,presence:optional"`
	UnknownValue      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfMDTNRExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfMDT-NR-ExtIEsEntry", v)
}

// DecodeAreaScopeOfMDTNRExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfMDTNRExtIEsEntry(data []byte) (AreaScopeOfMDTNRExtIEsEntry, int, error) {
	var v AreaScopeOfMDTNRExtIEsEntry
	err := ngapcodec.DecodeParameter("AreaScopeOfMDT-NR-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfMDTNRExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
