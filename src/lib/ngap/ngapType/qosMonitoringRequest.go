// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosMonitoringRequest represents ASN.1 QosMonitoringRequest.
const (
	QosMonitoringRequestPresentUl   Enumerated = 0
	QosMonitoringRequestPresentDl   Enumerated = 1
	QosMonitoringRequestPresentBoth Enumerated = 2
	QosMonitoringRequestPresentStop Enumerated = 3
)

type QosMonitoringRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosMonitoringRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosMonitoringRequest", v)
}

// DecodeQosMonitoringRequest decodes a standalone APER value and reports the consumed octets.
func DecodeQosMonitoringRequest(data []byte) (QosMonitoringRequest, int, error) {
	var v QosMonitoringRequest
	err := ngapcodec.DecodeParameter("QosMonitoringRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosMonitoringRequest{}, 0, err
	}
	return v, len(data), err
}
