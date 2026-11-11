// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceResumeItemRESRes represents ASN.1 PDUSessionResourceResumeItemRESRes.
type PDUSessionResourceResumeItemRESRes struct {
	PDUSessionID                    PDUSessionID
	UEContextResumeResponseTransfer []byte
	HasIEExtensions                 bool
	IEExtensions                    []PDUSessionResourceResumeItemRESResExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceResumeItemRESRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceResumeItemRESRes", v)
}

// DecodePDUSessionResourceResumeItemRESRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceResumeItemRESRes(data []byte) (PDUSessionResourceResumeItemRESRes, int, error) {
	var v PDUSessionResourceResumeItemRESRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceResumeItemRESRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceResumeItemRESRes{}, 0, err
	}
	return v, len(data), err
}
