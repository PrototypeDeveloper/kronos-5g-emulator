// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSwitchedList represents ASN.1 PDUSessionResourceSwitchedList.
type PDUSessionResourceSwitchedList struct {
	List []PDUSessionResourceSwitchedItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSwitchedList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSwitchedList", v)
}

// DecodePDUSessionResourceSwitchedList decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSwitchedList(data []byte) (PDUSessionResourceSwitchedList, int, error) {
	var v PDUSessionResourceSwitchedList
	err := ngapcodec.DecodeParameter("PDUSessionResourceSwitchedList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSwitchedList{}, 0, err
	}
	return v, len(data), err
}
