// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UpdateFeedback represents ASN.1 UpdateFeedback.
type UpdateFeedback struct {
	Value BitString `aper:"sizeExt,sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UpdateFeedback) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UpdateFeedback", v)
}

// DecodeUpdateFeedback decodes a standalone APER value and reports the consumed octets.
func DecodeUpdateFeedback(data []byte) (UpdateFeedback, int, error) {
	var v UpdateFeedback
	err := ngapcodec.DecodeParameter("UpdateFeedback", data, &v)
	if ngapcodec.IsFatal(err) {
		return UpdateFeedback{}, 0, err
	}
	return v, len(data), err
}
