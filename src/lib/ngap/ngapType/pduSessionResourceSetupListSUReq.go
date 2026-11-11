// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupListSUReq represents ASN.1 PDUSessionResourceSetupListSUReq.
type PDUSessionResourceSetupListSUReq struct {
	List []PDUSessionResourceSetupItemSUReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupListSUReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupListSUReq", v)
}

// DecodePDUSessionResourceSetupListSUReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupListSUReq(data []byte) (PDUSessionResourceSetupListSUReq, int, error) {
	var v PDUSessionResourceSetupListSUReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupListSUReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupListSUReq{}, 0, err
	}
	return v, len(data), err
}
