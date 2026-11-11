// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceNotifyItem represents ASN.1 PDUSessionResourceNotifyItem.
type PDUSessionResourceNotifyItem struct {
	PDUSessionID                     PDUSessionID
	PDUSessionResourceNotifyTransfer []byte
	HasIEExtensions                  bool
	IEExtensions                     []PDUSessionResourceNotifyItemExtIEsEntry
	UnknownExtensions                []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceNotifyItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceNotifyItem", v)
}

// DecodePDUSessionResourceNotifyItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceNotifyItem(data []byte) (PDUSessionResourceNotifyItem, int, error) {
	var v PDUSessionResourceNotifyItem
	err := ngapcodec.DecodeParameter("PDUSessionResourceNotifyItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceNotifyItem{}, 0, err
	}
	return v, len(data), err
}
