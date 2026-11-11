// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyItemModReq represents ASN.1 PDUSessionResourceModifyItemModReq.
type PDUSessionResourceModifyItemModReq struct {
	PDUSessionID                            PDUSessionID
	HasNASPDU                               bool
	NASPDU                                  NASPDU
	PDUSessionResourceModifyRequestTransfer []byte
	HasIEExtensions                         bool
	IEExtensions                            []PDUSessionResourceModifyItemModReqExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyItemModReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyItemModReq", v)
}

// DecodePDUSessionResourceModifyItemModReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyItemModReq(data []byte) (PDUSessionResourceModifyItemModReq, int, error) {
	var v PDUSessionResourceModifyItemModReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyItemModReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyItemModReq{}, 0, err
	}
	return v, len(data), err
}
