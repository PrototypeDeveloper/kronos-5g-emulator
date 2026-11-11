// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRUERLFReportContainer represents ASN.1 NRUERLFReportContainer.
type NRUERLFReportContainer struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRUERLFReportContainer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRUERLFReportContainer", v)
}

// DecodeNRUERLFReportContainer decodes a standalone APER value and reports the consumed octets.
func DecodeNRUERLFReportContainer(data []byte) (NRUERLFReportContainer, int, error) {
	var v NRUERLFReportContainer
	err := ngapcodec.DecodeParameter("NRUERLFReportContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRUERLFReportContainer{}, 0, err
	}
	return v, len(data), err
}
