// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupItemHOReq represents ASN.1 PDUSessionResourceSetupItemHOReq.
type PDUSessionResourceSetupItemHOReq struct {
	PDUSessionID            PDUSessionID
	SNSSAI                  SNSSAI
	HandoverRequestTransfer []byte
	HasIEExtensions         bool
	IEExtensions            []PDUSessionResourceSetupItemHOReqExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemHOReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemHOReq", v)
}

// DecodePDUSessionResourceSetupItemHOReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemHOReq(data []byte) (PDUSessionResourceSetupItemHOReq, int, error) {
	var v PDUSessionResourceSetupItemHOReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemHOReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemHOReq{}, 0, err
	}
	return v, len(data), err
}
