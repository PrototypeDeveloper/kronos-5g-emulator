// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionListMTCommHReqItem represents ASN.1 PDUSessionListMTCommHReq-Item.
type PDUSessionListMTCommHReqItem struct {
	PDUSessionID      PDUSessionID
	HasIEExtensions   bool
	IEExtensions      []PDUSessionListMTCommHReqItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionListMTCommHReqItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionListMTCommHReq-Item", v)
}

// DecodePDUSessionListMTCommHReqItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionListMTCommHReqItem(data []byte) (PDUSessionListMTCommHReqItem, int, error) {
	var v PDUSessionListMTCommHReqItem
	err := ngapcodec.DecodeParameter("PDUSessionListMTCommHReq-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionListMTCommHReqItem{}, 0, err
	}
	return v, len(data), err
}
