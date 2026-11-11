// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceItemCxtRelReq represents ASN.1 PDUSessionResourceItemCxtRelReq.
type PDUSessionResourceItemCxtRelReq struct {
	PDUSessionID      PDUSessionID
	HasIEExtensions   bool
	IEExtensions      []PDUSessionResourceItemCxtRelReqExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceItemCxtRelReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceItemCxtRelReq", v)
}

// DecodePDUSessionResourceItemCxtRelReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceItemCxtRelReq(data []byte) (PDUSessionResourceItemCxtRelReq, int, error) {
	var v PDUSessionResourceItemCxtRelReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceItemCxtRelReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceItemCxtRelReq{}, 0, err
	}
	return v, len(data), err
}
