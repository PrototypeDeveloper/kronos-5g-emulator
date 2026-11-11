// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DLNGUTNLInformationReused represents ASN.1 DL-NGU-TNLInformationReused.
const (
	DLNGUTNLInformationReusedPresentTrue Enumerated = 0
)

type DLNGUTNLInformationReused struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLNGUTNLInformationReused) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DL-NGU-TNLInformationReused", v)
}

// DecodeDLNGUTNLInformationReused decodes a standalone APER value and reports the consumed octets.
func DecodeDLNGUTNLInformationReused(data []byte) (DLNGUTNLInformationReused, int, error) {
	var v DLNGUTNLInformationReused
	err := ngapcodec.DecodeParameter("DL-NGU-TNLInformationReused", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLNGUTNLInformationReused{}, 0, err
	}
	return v, len(data), err
}
