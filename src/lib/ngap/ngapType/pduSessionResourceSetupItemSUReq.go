// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupItemSUReq represents ASN.1 PDUSessionResourceSetupItemSUReq.
type PDUSessionResourceSetupItemSUReq struct {
	PDUSessionID                           PDUSessionID
	HasPDUSessionNASPDU                    bool
	PDUSessionNASPDU                       NASPDU
	SNSSAI                                 SNSSAI
	PDUSessionResourceSetupRequestTransfer []byte
	HasIEExtensions                        bool
	IEExtensions                           []PDUSessionResourceSetupItemSUReqExtIEsEntry
	UnknownExtensions                      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemSUReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemSUReq", v)
}

// DecodePDUSessionResourceSetupItemSUReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemSUReq(data []byte) (PDUSessionResourceSetupItemSUReq, int, error) {
	var v PDUSessionResourceSetupItemSUReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemSUReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemSUReq{}, 0, err
	}
	return v, len(data), err
}
