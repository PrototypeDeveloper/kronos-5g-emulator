// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleasedListNot represents ASN.1 PDUSessionResourceReleasedListNot.
type PDUSessionResourceReleasedListNot struct {
	List []PDUSessionResourceReleasedItemNot `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedListNot) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedListNot", v)
}

// DecodePDUSessionResourceReleasedListNot decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedListNot(data []byte) (PDUSessionResourceReleasedListNot, int, error) {
	var v PDUSessionResourceReleasedListNot
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedListNot", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedListNot{}, 0, err
	}
	return v, len(data), err
}
