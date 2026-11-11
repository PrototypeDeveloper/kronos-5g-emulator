// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CEmodeBSupportIndicator represents ASN.1 CEmodeBSupport-Indicator.
const (
	CEmodeBSupportIndicatorPresentSupported Enumerated = 0
)

type CEmodeBSupportIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CEmodeBSupportIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CEmodeBSupport-Indicator", v)
}

// DecodeCEmodeBSupportIndicator decodes a standalone APER value and reports the consumed octets.
func DecodeCEmodeBSupportIndicator(data []byte) (CEmodeBSupportIndicator, int, error) {
	var v CEmodeBSupportIndicator
	err := ngapcodec.DecodeParameter("CEmodeBSupport-Indicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return CEmodeBSupportIndicator{}, 0, err
	}
	return v, len(data), err
}
