// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FailureIndication represents ASN.1 FailureIndication.
type FailureIndication struct {
	UERLFReportContainer UERLFReportContainer
	HasIEExtensions      bool
	IEExtensions         []FailureIndicationExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FailureIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FailureIndication", v)
}

// DecodeFailureIndication decodes a standalone APER value and reports the consumed octets.
func DecodeFailureIndication(data []byte) (FailureIndication, int, error) {
	var v FailureIndication
	err := ngapcodec.DecodeParameter("FailureIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return FailureIndication{}, 0, err
	}
	return v, len(data), err
}
