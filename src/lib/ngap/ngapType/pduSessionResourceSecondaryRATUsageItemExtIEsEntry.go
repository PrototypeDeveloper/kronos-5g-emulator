// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSecondaryRATUsageItemExtIEsEntryPresent uint32

const (
	PDUSessionResourceSecondaryRATUsageItemExtIEsEntryPresentNothing PDUSessionResourceSecondaryRATUsageItemExtIEsEntryPresent = iota
	PDUSessionResourceSecondaryRATUsageItemExtIEsEntryPresentUnknown
)

type PDUSessionResourceSecondaryRATUsageItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceSecondaryRATUsageItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSecondaryRATUsageItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSecondaryRATUsageItem-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSecondaryRATUsageItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSecondaryRATUsageItemExtIEsEntry(data []byte) (PDUSessionResourceSecondaryRATUsageItemExtIEsEntry, int, error) {
	var v PDUSessionResourceSecondaryRATUsageItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSecondaryRATUsageItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSecondaryRATUsageItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
