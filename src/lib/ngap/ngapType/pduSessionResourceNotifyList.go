// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceNotifyList represents ASN.1 PDUSessionResourceNotifyList.
type PDUSessionResourceNotifyList struct {
	List []PDUSessionResourceNotifyItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceNotifyList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceNotifyList", v)
}

// DecodePDUSessionResourceNotifyList decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceNotifyList(data []byte) (PDUSessionResourceNotifyList, int, error) {
	var v PDUSessionResourceNotifyList
	err := ngapcodec.DecodeParameter("PDUSessionResourceNotifyList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceNotifyList{}, 0, err
	}
	return v, len(data), err
}
