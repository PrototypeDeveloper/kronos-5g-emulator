// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceAdmittedList represents ASN.1 PDUSessionResourceAdmittedList.
type PDUSessionResourceAdmittedList struct {
	List []PDUSessionResourceAdmittedItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceAdmittedList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceAdmittedList", v)
}

// DecodePDUSessionResourceAdmittedList decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceAdmittedList(data []byte) (PDUSessionResourceAdmittedList, int, error) {
	var v PDUSessionResourceAdmittedList
	err := ngapcodec.DecodeParameter("PDUSessionResourceAdmittedList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceAdmittedList{}, 0, err
	}
	return v, len(data), err
}
