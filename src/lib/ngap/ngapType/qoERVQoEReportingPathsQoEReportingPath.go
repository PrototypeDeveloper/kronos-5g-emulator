// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoERVQoEReportingPathsQoEReportingPath represents ASN.1 QoERVQoEReportingPathsQoEReportingPath.
const (
	QoERVQoEReportingPathsQoEReportingPathPresentSrb4 Enumerated = 0
	QoERVQoEReportingPathsQoEReportingPathPresentSrb5 Enumerated = 1
)

type QoERVQoEReportingPathsQoEReportingPath struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v QoERVQoEReportingPathsQoEReportingPath) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:QoERVQoEReportingPathsQoEReportingPath", v)
}

// DecodeQoERVQoEReportingPathsQoEReportingPath decodes a standalone APER value and reports consumed octets.
func DecodeQoERVQoEReportingPathsQoEReportingPath(data []byte) (QoERVQoEReportingPathsQoEReportingPath, int, error) {
	var v QoERVQoEReportingPathsQoEReportingPath
	err := ngapcodec.DecodeParameter("inline:QoERVQoEReportingPathsQoEReportingPath", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoERVQoEReportingPathsQoEReportingPath{}, 0, err
	}
	return v, len(data), err
}
