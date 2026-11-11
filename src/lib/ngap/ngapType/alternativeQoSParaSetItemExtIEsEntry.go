// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AlternativeQoSParaSetItemExtIEsEntryPresent uint32

const (
	AlternativeQoSParaSetItemExtIEsEntryPresentNothing AlternativeQoSParaSetItemExtIEsEntryPresent = iota
	AlternativeQoSParaSetItemExtIEsEntryPresentMaximumDataBurstVolume
	AlternativeQoSParaSetItemExtIEsEntryPresentPduSetDelayBudgetDownlink
	AlternativeQoSParaSetItemExtIEsEntryPresentPduSetDelayBudgetUplink
	AlternativeQoSParaSetItemExtIEsEntryPresentPduSetErrorRateDownlink
	AlternativeQoSParaSetItemExtIEsEntryPresentPduSetErrorRateUplink
	AlternativeQoSParaSetItemExtIEsEntryPresentUnknown
)

type AlternativeQoSParaSetItemExtIEsEntry struct {
	ID                        ProtocolExtensionID
	Criticality               Criticality
	Present                   AlternativeQoSParaSetItemExtIEsEntryPresent
	MaximumDataBurstVolume    MaximumDataBurstVolume    `ngap:"id:432,criticality:ignore,presence:optional"`
	PduSetDelayBudgetDownlink ExtendedPacketDelayBudget `ngap:"id:484,criticality:ignore,presence:optional"`
	PduSetDelayBudgetUplink   ExtendedPacketDelayBudget `ngap:"id:485,criticality:ignore,presence:optional"`
	PduSetErrorRateDownlink   PacketErrorRate           `ngap:"id:486,criticality:ignore,presence:optional"`
	PduSetErrorRateUplink     PacketErrorRate           `ngap:"id:487,criticality:ignore,presence:optional"`
	UnknownValue              []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AlternativeQoSParaSetItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AlternativeQoSParaSetItem-ExtIEsEntry", v)
}

// DecodeAlternativeQoSParaSetItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAlternativeQoSParaSetItemExtIEsEntry(data []byte) (AlternativeQoSParaSetItemExtIEsEntry, int, error) {
	var v AlternativeQoSParaSetItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AlternativeQoSParaSetItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AlternativeQoSParaSetItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
