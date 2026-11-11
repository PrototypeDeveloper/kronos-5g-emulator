// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSupportIndicator represents ASN.1 MBS-SupportIndicator.
const (
	MBSSupportIndicatorPresentMulticastSupported                           Enumerated = 0
	MBSSupportIndicatorPresentMulticastSupportedWithReceptionInRRCInactive Enumerated = 1
)

type MBSSupportIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSupportIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SupportIndicator", v)
}

// DecodeMBSSupportIndicator decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSupportIndicator(data []byte) (MBSSupportIndicator, int, error) {
	var v MBSSupportIndicator
	err := ngapcodec.DecodeParameter("MBS-SupportIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSupportIndicator{}, 0, err
	}
	return v, len(data), err
}
