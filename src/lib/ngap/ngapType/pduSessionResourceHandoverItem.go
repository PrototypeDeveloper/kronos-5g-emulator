// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceHandoverItem represents ASN.1 PDUSessionResourceHandoverItem.
type PDUSessionResourceHandoverItem struct {
	PDUSessionID            PDUSessionID
	HandoverCommandTransfer []byte
	HasIEExtensions         bool
	IEExtensions            []PDUSessionResourceHandoverItemExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceHandoverItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceHandoverItem", v)
}

// DecodePDUSessionResourceHandoverItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceHandoverItem(data []byte) (PDUSessionResourceHandoverItem, int, error) {
	var v PDUSessionResourceHandoverItem
	err := ngapcodec.DecodeParameter("PDUSessionResourceHandoverItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceHandoverItem{}, 0, err
	}
	return v, len(data), err
}
