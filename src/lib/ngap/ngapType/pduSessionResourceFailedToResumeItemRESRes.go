// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToResumeItemRESRes represents ASN.1 PDUSessionResourceFailedToResumeItemRESRes.
type PDUSessionResourceFailedToResumeItemRESRes struct {
	PDUSessionID      PDUSessionID
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []PDUSessionResourceFailedToResumeItemRESResExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToResumeItemRESRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToResumeItemRESRes", v)
}

// DecodePDUSessionResourceFailedToResumeItemRESRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToResumeItemRESRes(data []byte) (PDUSessionResourceFailedToResumeItemRESRes, int, error) {
	var v PDUSessionResourceFailedToResumeItemRESRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToResumeItemRESRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToResumeItemRESRes{}, 0, err
	}
	return v, len(data), err
}
