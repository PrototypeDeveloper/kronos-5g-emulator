// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupListPSReq represents ASN.1 PDUSessionResourceFailedToSetupListPSReq.
type PDUSessionResourceFailedToSetupListPSReq struct {
	List []PDUSessionResourceFailedToSetupItemPSReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupListPSReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupListPSReq", v)
}

// DecodePDUSessionResourceFailedToSetupListPSReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupListPSReq(data []byte) (PDUSessionResourceFailedToSetupListPSReq, int, error) {
	var v PDUSessionResourceFailedToSetupListPSReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupListPSReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupListPSReq{}, 0, err
	}
	return v, len(data), err
}
