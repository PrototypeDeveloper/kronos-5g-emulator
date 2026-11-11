// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupItemCxtReq represents ASN.1 PDUSessionResourceSetupItemCxtReq.
type PDUSessionResourceSetupItemCxtReq struct {
	PDUSessionID                           PDUSessionID
	HasNASPDU                              bool
	NASPDU                                 NASPDU
	SNSSAI                                 SNSSAI
	PDUSessionResourceSetupRequestTransfer []byte
	HasIEExtensions                        bool
	IEExtensions                           []PDUSessionResourceSetupItemCxtReqExtIEsEntry
	UnknownExtensions                      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemCxtReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemCxtReq", v)
}

// DecodePDUSessionResourceSetupItemCxtReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemCxtReq(data []byte) (PDUSessionResourceSetupItemCxtReq, int, error) {
	var v PDUSessionResourceSetupItemCxtReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemCxtReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemCxtReq{}, 0, err
	}
	return v, len(data), err
}
