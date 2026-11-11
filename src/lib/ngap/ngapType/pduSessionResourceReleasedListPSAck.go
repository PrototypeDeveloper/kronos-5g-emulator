// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleasedListPSAck represents ASN.1 PDUSessionResourceReleasedListPSAck.
type PDUSessionResourceReleasedListPSAck struct {
	List []PDUSessionResourceReleasedItemPSAck `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedListPSAck) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedListPSAck", v)
}

// DecodePDUSessionResourceReleasedListPSAck decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedListPSAck(data []byte) (PDUSessionResourceReleasedListPSAck, int, error) {
	var v PDUSessionResourceReleasedListPSAck
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedListPSAck", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedListPSAck{}, 0, err
	}
	return v, len(data), err
}
