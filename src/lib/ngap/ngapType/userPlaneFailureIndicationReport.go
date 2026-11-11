// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UserPlaneFailureIndicationReport represents ASN.1 UserPlaneFailureIndicationReport.
const (
	UserPlaneFailureIndicationReportPresentQosFlowsAndTunnelToBeReleased Enumerated = 0
	UserPlaneFailureIndicationReportPresentQoSFlowsToBeMoved             Enumerated = 1
	UserPlaneFailureIndicationReportPresentNewTransportAddressAllocated  Enumerated = 2
)

type UserPlaneFailureIndicationReport struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserPlaneFailureIndicationReport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserPlaneFailureIndicationReport", v)
}

// DecodeUserPlaneFailureIndicationReport decodes a standalone APER value and reports the consumed octets.
func DecodeUserPlaneFailureIndicationReport(data []byte) (UserPlaneFailureIndicationReport, int, error) {
	var v UserPlaneFailureIndicationReport
	err := ngapcodec.DecodeParameter("UserPlaneFailureIndicationReport", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserPlaneFailureIndicationReport{}, 0, err
	}
	return v, len(data), err
}
