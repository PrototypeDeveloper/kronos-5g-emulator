// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IMSVoiceSupportIndicator represents ASN.1 IMSVoiceSupportIndicator.
const (
	IMSVoiceSupportIndicatorPresentSupported    Enumerated = 0
	IMSVoiceSupportIndicatorPresentNotSupported Enumerated = 1
)

type IMSVoiceSupportIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IMSVoiceSupportIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IMSVoiceSupportIndicator", v)
}

// DecodeIMSVoiceSupportIndicator decodes a standalone APER value and reports the consumed octets.
func DecodeIMSVoiceSupportIndicator(data []byte) (IMSVoiceSupportIndicator, int, error) {
	var v IMSVoiceSupportIndicator
	err := ngapcodec.DecodeParameter("IMSVoiceSupportIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return IMSVoiceSupportIndicator{}, 0, err
	}
	return v, len(data), err
}
