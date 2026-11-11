// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleasedListPSFail represents ASN.1 PDUSessionResourceReleasedListPSFail.
type PDUSessionResourceReleasedListPSFail struct {
	List []PDUSessionResourceReleasedItemPSFail `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedListPSFail) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedListPSFail", v)
}

// DecodePDUSessionResourceReleasedListPSFail decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedListPSFail(data []byte) (PDUSessionResourceReleasedListPSFail, int, error) {
	var v PDUSessionResourceReleasedListPSFail
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedListPSFail", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedListPSFail{}, 0, err
	}
	return v, len(data), err
}
