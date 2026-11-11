// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoERVQoEReportingPathsRVQoEReportingPath represents ASN.1 QoERVQoEReportingPathsRVQoEReportingPath.
const (
	QoERVQoEReportingPathsRVQoEReportingPathPresentSrb4 Enumerated = 0
	QoERVQoEReportingPathsRVQoEReportingPathPresentSrb5 Enumerated = 1
)

type QoERVQoEReportingPathsRVQoEReportingPath struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v QoERVQoEReportingPathsRVQoEReportingPath) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:QoERVQoEReportingPathsRVQoEReportingPath", v)
}

// DecodeQoERVQoEReportingPathsRVQoEReportingPath decodes a standalone APER value and reports consumed octets.
func DecodeQoERVQoEReportingPathsRVQoEReportingPath(data []byte) (QoERVQoEReportingPathsRVQoEReportingPath, int, error) {
	var v QoERVQoEReportingPathsRVQoEReportingPath
	err := ngapcodec.DecodeParameter("inline:QoERVQoEReportingPathsRVQoEReportingPath", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoERVQoEReportingPathsRVQoEReportingPath{}, 0, err
	}
	return v, len(data), err
}
