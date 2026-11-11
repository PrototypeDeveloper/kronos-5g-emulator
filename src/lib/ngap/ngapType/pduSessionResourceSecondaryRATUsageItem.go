// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSecondaryRATUsageItem represents ASN.1 PDUSessionResourceSecondaryRATUsageItem.
type PDUSessionResourceSecondaryRATUsageItem struct {
	PDUSessionID                        PDUSessionID
	SecondaryRATDataUsageReportTransfer []byte
	HasIEExtensions                     bool
	IEExtensions                        []PDUSessionResourceSecondaryRATUsageItemExtIEsEntry
	UnknownExtensions                   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSecondaryRATUsageItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSecondaryRATUsageItem", v)
}

// DecodePDUSessionResourceSecondaryRATUsageItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSecondaryRATUsageItem(data []byte) (PDUSessionResourceSecondaryRATUsageItem, int, error) {
	var v PDUSessionResourceSecondaryRATUsageItem
	err := ngapcodec.DecodeParameter("PDUSessionResourceSecondaryRATUsageItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSecondaryRATUsageItem{}, 0, err
	}
	return v, len(data), err
}
