// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// N6JitterInformation represents ASN.1 N6JitterInformation.
type N6JitterInformation struct {
	N6JitterLowerBound int64 `aper:"valueLB:-127,valueUB:127"`
	N6JitterUpperBound int64 `aper:"valueLB:-127,valueUB:127"`
	HasIEExtensions    bool
	IEExtensions       []N6JitterInformationExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v N6JitterInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("N6JitterInformation", v)
}

// DecodeN6JitterInformation decodes a standalone APER value and reports the consumed octets.
func DecodeN6JitterInformation(data []byte) (N6JitterInformation, int, error) {
	var v N6JitterInformation
	err := ngapcodec.DecodeParameter("N6JitterInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return N6JitterInformation{}, 0, err
	}
	return v, len(data), err
}
