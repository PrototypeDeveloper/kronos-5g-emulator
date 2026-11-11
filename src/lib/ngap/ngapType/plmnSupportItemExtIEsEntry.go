// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PLMNSupportItemExtIEsEntryPresent uint32

const (
	PLMNSupportItemExtIEsEntryPresentNothing PLMNSupportItemExtIEsEntryPresent = iota
	PLMNSupportItemExtIEsEntryPresentNPNSupport
	PLMNSupportItemExtIEsEntryPresentExtendedSliceSupportList
	PLMNSupportItemExtIEsEntryPresentOnboardingSupport
	PLMNSupportItemExtIEsEntryPresentUnknown
)

type PLMNSupportItemExtIEsEntry struct {
	ID                       ProtocolExtensionID
	Criticality              Criticality
	Present                  PLMNSupportItemExtIEsEntryPresent
	NPNSupport               NPNSupport               `ngap:"id:258,criticality:reject,presence:optional"`
	ExtendedSliceSupportList ExtendedSliceSupportList `ngap:"id:270,criticality:reject,presence:optional"`
	OnboardingSupport        OnboardingSupport        `ngap:"id:325,criticality:ignore,presence:optional"`
	UnknownValue             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PLMNSupportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PLMNSupportItem-ExtIEsEntry", v)
}

// DecodePLMNSupportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePLMNSupportItemExtIEsEntry(data []byte) (PLMNSupportItemExtIEsEntry, int, error) {
	var v PLMNSupportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PLMNSupportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PLMNSupportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
