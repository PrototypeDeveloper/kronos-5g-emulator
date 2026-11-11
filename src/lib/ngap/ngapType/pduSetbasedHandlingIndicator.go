// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSetbasedHandlingIndicator represents ASN.1 PDUSetbasedHandlingIndicator.
const (
	PDUSetbasedHandlingIndicatorPresentSupported Enumerated = 0
)

type PDUSetbasedHandlingIndicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSetbasedHandlingIndicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSetbasedHandlingIndicator", v)
}

// DecodePDUSetbasedHandlingIndicator decodes a standalone APER value and reports the consumed octets.
func DecodePDUSetbasedHandlingIndicator(data []byte) (PDUSetbasedHandlingIndicator, int, error) {
	var v PDUSetbasedHandlingIndicator
	err := ngapcodec.DecodeParameter("PDUSetbasedHandlingIndicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSetbasedHandlingIndicator{}, 0, err
	}
	return v, len(data), err
}
