// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ULNGUUPTNLModifyItemExtIEsEntryPresent uint32

const (
	ULNGUUPTNLModifyItemExtIEsEntryPresentNothing ULNGUUPTNLModifyItemExtIEsEntryPresent = iota
	ULNGUUPTNLModifyItemExtIEsEntryPresentRedundantULNGUUPTNLInformation
	ULNGUUPTNLModifyItemExtIEsEntryPresentRedundantDLNGUUPTNLInformation
	ULNGUUPTNLModifyItemExtIEsEntryPresentUnknown
)

type ULNGUUPTNLModifyItemExtIEsEntry struct {
	ID                             ProtocolExtensionID
	Criticality                    Criticality
	Present                        ULNGUUPTNLModifyItemExtIEsEntryPresent
	RedundantULNGUUPTNLInformation UPTransportLayerInformation `ngap:"id:195,criticality:ignore,presence:optional"`
	RedundantDLNGUUPTNLInformation UPTransportLayerInformation `ngap:"id:192,criticality:ignore,presence:optional"`
	UnknownValue                   []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ULNGUUPTNLModifyItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UL-NGU-UP-TNLModifyItem-ExtIEsEntry", v)
}

// DecodeULNGUUPTNLModifyItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeULNGUUPTNLModifyItemExtIEsEntry(data []byte) (ULNGUUPTNLModifyItemExtIEsEntry, int, error) {
	var v ULNGUUPTNLModifyItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UL-NGU-UP-TNLModifyItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ULNGUUPTNLModifyItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
