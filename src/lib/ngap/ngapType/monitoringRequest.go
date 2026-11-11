// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MonitoringRequest represents ASN.1 MonitoringRequest.
const (
	MonitoringRequestPresentUl   Enumerated = 0
	MonitoringRequestPresentDl   Enumerated = 1
	MonitoringRequestPresentBoth Enumerated = 2
	MonitoringRequestPresentStop Enumerated = 3
)

type MonitoringRequest struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MonitoringRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MonitoringRequest", v)
}

// DecodeMonitoringRequest decodes a standalone APER value and reports the consumed octets.
func DecodeMonitoringRequest(data []byte) (MonitoringRequest, int, error) {
	var v MonitoringRequest
	err := ngapcodec.DecodeParameter("MonitoringRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return MonitoringRequest{}, 0, err
	}
	return v, len(data), err
}
