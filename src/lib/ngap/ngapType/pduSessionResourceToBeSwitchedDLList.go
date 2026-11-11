// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceToBeSwitchedDLList represents ASN.1 PDUSessionResourceToBeSwitchedDLList.
type PDUSessionResourceToBeSwitchedDLList struct {
	List []PDUSessionResourceToBeSwitchedDLItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToBeSwitchedDLList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToBeSwitchedDLList", v)
}

// DecodePDUSessionResourceToBeSwitchedDLList decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToBeSwitchedDLList(data []byte) (PDUSessionResourceToBeSwitchedDLList, int, error) {
	var v PDUSessionResourceToBeSwitchedDLList
	err := ngapcodec.DecodeParameter("PDUSessionResourceToBeSwitchedDLList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToBeSwitchedDLList{}, 0, err
	}
	return v, len(data), err
}
