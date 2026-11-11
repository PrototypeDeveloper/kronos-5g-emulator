// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceResumeItemRESReq represents ASN.1 PDUSessionResourceResumeItemRESReq.
type PDUSessionResourceResumeItemRESReq struct {
	PDUSessionID                   PDUSessionID
	UEContextResumeRequestTransfer []byte
	HasIEExtensions                bool
	IEExtensions                   []PDUSessionResourceResumeItemRESReqExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceResumeItemRESReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceResumeItemRESReq", v)
}

// DecodePDUSessionResourceResumeItemRESReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceResumeItemRESReq(data []byte) (PDUSessionResourceResumeItemRESReq, int, error) {
	var v PDUSessionResourceResumeItemRESReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceResumeItemRESReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceResumeItemRESReq{}, 0, err
	}
	return v, len(data), err
}
