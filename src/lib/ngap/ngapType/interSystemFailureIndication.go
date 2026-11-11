// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InterSystemFailureIndication represents ASN.1 InterSystemFailureIndication.
type InterSystemFailureIndication struct {
	HasUERLFReportContainer bool
	UERLFReportContainer    UERLFReportContainer
	HasIEExtensions         bool
	IEExtensions            []InterSystemFailureIndicationExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InterSystemFailureIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InterSystemFailureIndication", v)
}

// DecodeInterSystemFailureIndication decodes a standalone APER value and reports the consumed octets.
func DecodeInterSystemFailureIndication(data []byte) (InterSystemFailureIndication, int, error) {
	var v InterSystemFailureIndication
	err := ngapcodec.DecodeParameter("InterSystemFailureIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return InterSystemFailureIndication{}, 0, err
	}
	return v, len(data), err
}
