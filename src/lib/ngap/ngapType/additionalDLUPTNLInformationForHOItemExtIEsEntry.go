// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AdditionalDLUPTNLInformationForHOItemExtIEsEntryPresent uint32

const (
	AdditionalDLUPTNLInformationForHOItemExtIEsEntryPresentNothing AdditionalDLUPTNLInformationForHOItemExtIEsEntryPresent = iota
	AdditionalDLUPTNLInformationForHOItemExtIEsEntryPresentAdditionalRedundantDLNGUUPTNLInformation
	AdditionalDLUPTNLInformationForHOItemExtIEsEntryPresentUnknown
)

type AdditionalDLUPTNLInformationForHOItemExtIEsEntry struct {
	ID                                       ProtocolExtensionID
	Criticality                              Criticality
	Present                                  AdditionalDLUPTNLInformationForHOItemExtIEsEntryPresent
	AdditionalRedundantDLNGUUPTNLInformation UPTransportLayerInformation `ngap:"id:183,criticality:ignore,presence:optional"`
	UnknownValue                             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalDLUPTNLInformationForHOItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AdditionalDLUPTNLInformationForHOItem-ExtIEsEntry", v)
}

// DecodeAdditionalDLUPTNLInformationForHOItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalDLUPTNLInformationForHOItemExtIEsEntry(data []byte) (AdditionalDLUPTNLInformationForHOItemExtIEsEntry, int, error) {
	var v AdditionalDLUPTNLInformationForHOItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AdditionalDLUPTNLInformationForHOItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalDLUPTNLInformationForHOItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
