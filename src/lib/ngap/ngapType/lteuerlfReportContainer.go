// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LTEUERLFReportContainer represents ASN.1 LTEUERLFReportContainer.
type LTEUERLFReportContainer struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LTEUERLFReportContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LTEUERLFReportContainer", v)
}

// DecodeLTEUERLFReportContainer decodes a standalone APER value and reports the consumed octets.
func DecodeLTEUERLFReportContainer(data []byte) (LTEUERLFReportContainer, int, error) {
	var v LTEUERLFReportContainer
	err := ngapcodec.DecodeParameter("LTEUERLFReportContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEUERLFReportContainer{}, 0, err
	}
	return v, len(data), err
}
