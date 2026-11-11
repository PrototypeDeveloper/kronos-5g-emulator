// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupListHOAck represents ASN.1 PDUSessionResourceFailedToSetupListHOAck.
type PDUSessionResourceFailedToSetupListHOAck struct {
	List []PDUSessionResourceFailedToSetupItemHOAck `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupListHOAck) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupListHOAck", v)
}

// DecodePDUSessionResourceFailedToSetupListHOAck decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupListHOAck(data []byte) (PDUSessionResourceFailedToSetupListHOAck, int, error) {
	var v PDUSessionResourceFailedToSetupListHOAck
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupListHOAck", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupListHOAck{}, 0, err
	}
	return v, len(data), err
}
