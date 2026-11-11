// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANfeedbacktypeReactive represents ASN.1 RANfeedbacktype-reactive.
type RANfeedbacktypeReactive struct {
	CapabilityForBATAdaptation RANfeedbacktypeReactiveCapabilityForBATAdaptation `aper:"valueExt,valueLB:0,valueUB:0"`
	HasIEExtension             bool
	IEExtension                []RANfeedbacktypeReactiveExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANfeedbacktypeReactive) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANfeedbacktype-reactive", v)
}

// DecodeRANfeedbacktypeReactive decodes a standalone APER value and reports the consumed octets.
func DecodeRANfeedbacktypeReactive(data []byte) (RANfeedbacktypeReactive, int, error) {
	var v RANfeedbacktypeReactive
	err := ngapcodec.DecodeParameter("RANfeedbacktype-reactive", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANfeedbacktypeReactive{}, 0, err
	}
	return v, len(data), err
}
