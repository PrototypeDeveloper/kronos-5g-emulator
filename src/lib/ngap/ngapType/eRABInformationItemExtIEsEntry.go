// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ERABInformationItemExtIEsEntryPresent uint32

const (
	ERABInformationItemExtIEsEntryPresentNothing ERABInformationItemExtIEsEntryPresent = iota
	ERABInformationItemExtIEsEntryPresentSourceTNLAddrInfo
	ERABInformationItemExtIEsEntryPresentSourceNodeTNLAddrInfo
	ERABInformationItemExtIEsEntryPresentUnknown
)

type ERABInformationItemExtIEsEntry struct {
	ID                    ProtocolExtensionID
	Criticality           Criticality
	Present               ERABInformationItemExtIEsEntryPresent
	SourceTNLAddrInfo     TransportLayerAddress `ngap:"id:284,criticality:ignore,presence:optional"`
	SourceNodeTNLAddrInfo TransportLayerAddress `ngap:"id:354,criticality:ignore,presence:optional"`
	UnknownValue          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ERABInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("E-RABInformationItem-ExtIEsEntry", v)
}

// DecodeERABInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeERABInformationItemExtIEsEntry(data []byte) (ERABInformationItemExtIEsEntry, int, error) {
	var v ERABInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("E-RABInformationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ERABInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
