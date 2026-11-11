// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupListHOReq represents ASN.1 PDUSessionResourceSetupListHOReq.
type PDUSessionResourceSetupListHOReq struct {
	List []PDUSessionResourceSetupItemHOReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupListHOReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupListHOReq", v)
}

// DecodePDUSessionResourceSetupListHOReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupListHOReq(data []byte) (PDUSessionResourceSetupListHOReq, int, error) {
	var v PDUSessionResourceSetupListHOReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupListHOReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupListHOReq{}, 0, err
	}
	return v, len(data), err
}
