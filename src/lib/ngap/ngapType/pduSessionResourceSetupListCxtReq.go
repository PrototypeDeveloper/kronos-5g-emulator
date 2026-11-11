// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupListCxtReq represents ASN.1 PDUSessionResourceSetupListCxtReq.
type PDUSessionResourceSetupListCxtReq struct {
	List []PDUSessionResourceSetupItemCxtReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupListCxtReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupListCxtReq", v)
}

// DecodePDUSessionResourceSetupListCxtReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupListCxtReq(data []byte) (PDUSessionResourceSetupListCxtReq, int, error) {
	var v PDUSessionResourceSetupListCxtReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupListCxtReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupListCxtReq{}, 0, err
	}
	return v, len(data), err
}
