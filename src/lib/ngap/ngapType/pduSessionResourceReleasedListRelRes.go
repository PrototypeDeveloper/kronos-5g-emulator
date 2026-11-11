// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleasedListRelRes represents ASN.1 PDUSessionResourceReleasedListRelRes.
type PDUSessionResourceReleasedListRelRes struct {
	List []PDUSessionResourceReleasedItemRelRes `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedListRelRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedListRelRes", v)
}

// DecodePDUSessionResourceReleasedListRelRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedListRelRes(data []byte) (PDUSessionResourceReleasedListRelRes, int, error) {
	var v PDUSessionResourceReleasedListRelRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedListRelRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedListRelRes{}, 0, err
	}
	return v, len(data), err
}
