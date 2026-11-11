// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RRCInactiveTransitionReportRequest represents ASN.1 RRCInactiveTransitionReportRequest.
const (
	RRCInactiveTransitionReportRequestPresentSubsequentStateTransitionReport Enumerated = 0
	RRCInactiveTransitionReportRequestPresentSingleRrcConnectedStateReport   Enumerated = 1
	RRCInactiveTransitionReportRequestPresentCancelReport                    Enumerated = 2
)

type RRCInactiveTransitionReportRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RRCInactiveTransitionReportRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RRCInactiveTransitionReportRequest", v)
}

// DecodeRRCInactiveTransitionReportRequest decodes a standalone APER value and reports the consumed octets.
func DecodeRRCInactiveTransitionReportRequest(data []byte) (RRCInactiveTransitionReportRequest, int, error) {
	var v RRCInactiveTransitionReportRequest
	err := ngapcodec.DecodeParameter("RRCInactiveTransitionReportRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return RRCInactiveTransitionReportRequest{}, 0, err
	}
	return v, len(data), err
}
