// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BroadcastPLMNItemExtIEsEntryPresent uint32

const (
	BroadcastPLMNItemExtIEsEntryPresentNothing BroadcastPLMNItemExtIEsEntryPresent = iota
	BroadcastPLMNItemExtIEsEntryPresentNPNSupport
	BroadcastPLMNItemExtIEsEntryPresentExtendedTAISliceSupportList
	BroadcastPLMNItemExtIEsEntryPresentTAINSAGSupportList
	BroadcastPLMNItemExtIEsEntryPresentUnknown
)

type BroadcastPLMNItemExtIEsEntry struct {
	ID                          ProtocolExtensionID
	Criticality                 Criticality
	Present                     BroadcastPLMNItemExtIEsEntryPresent
	NPNSupport                  NPNSupport               `ngap:"id:258,criticality:reject,presence:optional"`
	ExtendedTAISliceSupportList ExtendedSliceSupportList `ngap:"id:271,criticality:reject,presence:optional"`
	TAINSAGSupportList          TAINSAGSupportList       `ngap:"id:353,criticality:ignore,presence:optional"`
	UnknownValue                []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastPLMNItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastPLMNItem-ExtIEsEntry", v)
}

// DecodeBroadcastPLMNItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastPLMNItemExtIEsEntry(data []byte) (BroadcastPLMNItemExtIEsEntry, int, error) {
	var v BroadcastPLMNItemExtIEsEntry
	err := ngapcodec.DecodeParameter("BroadcastPLMNItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastPLMNItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
