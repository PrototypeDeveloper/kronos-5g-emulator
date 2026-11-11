// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type Dynamic5QIDescriptorExtIEsEntryPresent uint32

const (
	Dynamic5QIDescriptorExtIEsEntryPresentNothing Dynamic5QIDescriptorExtIEsEntryPresent = iota
	Dynamic5QIDescriptorExtIEsEntryPresentExtendedPacketDelayBudget
	Dynamic5QIDescriptorExtIEsEntryPresentCNPacketDelayBudgetDL
	Dynamic5QIDescriptorExtIEsEntryPresentCNPacketDelayBudgetUL
	Dynamic5QIDescriptorExtIEsEntryPresentUnknown
)

type Dynamic5QIDescriptorExtIEsEntry struct {
	ID                        ProtocolExtensionID
	Criticality               Criticality
	Present                   Dynamic5QIDescriptorExtIEsEntryPresent
	ExtendedPacketDelayBudget ExtendedPacketDelayBudget `ngap:"id:189,criticality:ignore,presence:optional"`
	CNPacketDelayBudgetDL     ExtendedPacketDelayBudget `ngap:"id:187,criticality:ignore,presence:optional"`
	CNPacketDelayBudgetUL     ExtendedPacketDelayBudget `ngap:"id:188,criticality:ignore,presence:optional"`
	UnknownValue              []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Dynamic5QIDescriptorExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Dynamic5QIDescriptor-ExtIEsEntry", v)
}

// DecodeDynamic5QIDescriptorExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDynamic5QIDescriptorExtIEsEntry(data []byte) (Dynamic5QIDescriptorExtIEsEntry, int, error) {
	var v Dynamic5QIDescriptorExtIEsEntry
	err := ngapcodec.DecodeParameter("Dynamic5QIDescriptor-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return Dynamic5QIDescriptorExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
