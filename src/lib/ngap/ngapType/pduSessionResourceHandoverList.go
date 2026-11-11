// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceHandoverList represents ASN.1 PDUSessionResourceHandoverList.
type PDUSessionResourceHandoverList struct {
	List []PDUSessionResourceHandoverItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceHandoverList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceHandoverList", v)
}

// DecodePDUSessionResourceHandoverList decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceHandoverList(data []byte) (PDUSessionResourceHandoverList, int, error) {
	var v PDUSessionResourceHandoverList
	err := ngapcodec.DecodeParameter("PDUSessionResourceHandoverList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceHandoverList{}, 0, err
	}
	return v, len(data), err
}
