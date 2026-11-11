// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NonDynamic5QIDescriptorExtIEsEntryPresent uint32

const (
	NonDynamic5QIDescriptorExtIEsEntryPresentNothing NonDynamic5QIDescriptorExtIEsEntryPresent = iota
	NonDynamic5QIDescriptorExtIEsEntryPresentCNPacketDelayBudgetDL
	NonDynamic5QIDescriptorExtIEsEntryPresentCNPacketDelayBudgetUL
	NonDynamic5QIDescriptorExtIEsEntryPresentUnknown
)

type NonDynamic5QIDescriptorExtIEsEntry struct {
	ID                    ProtocolExtensionID
	Criticality           Criticality
	Present               NonDynamic5QIDescriptorExtIEsEntryPresent
	CNPacketDelayBudgetDL ExtendedPacketDelayBudget `ngap:"id:187,criticality:ignore,presence:optional"`
	CNPacketDelayBudgetUL ExtendedPacketDelayBudget `ngap:"id:188,criticality:ignore,presence:optional"`
	UnknownValue          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NonDynamic5QIDescriptorExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NonDynamic5QIDescriptor-ExtIEsEntry", v)
}

// DecodeNonDynamic5QIDescriptorExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNonDynamic5QIDescriptorExtIEsEntry(data []byte) (NonDynamic5QIDescriptorExtIEsEntry, int, error) {
	var v NonDynamic5QIDescriptorExtIEsEntry
	err := ngapcodec.DecodeParameter("NonDynamic5QIDescriptor-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NonDynamic5QIDescriptorExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
