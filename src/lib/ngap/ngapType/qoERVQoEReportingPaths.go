// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoERVQoEReportingPaths represents ASN.1 QoERVQoEReportingPaths.
type QoERVQoEReportingPaths struct {
	HasQoEReportingPath   bool
	QoEReportingPath      QoERVQoEReportingPathsQoEReportingPath `aper:"valueExt,valueLB:0,valueUB:1"`
	HasRVQoEReportingPath bool
	RVQoEReportingPath    QoERVQoEReportingPathsRVQoEReportingPath `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions       bool
	IEExtensions          []QoERVQoEReportingPathsExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoERVQoEReportingPaths) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoERVQoEReportingPaths", v)
}

// DecodeQoERVQoEReportingPaths decodes a standalone APER value and reports the consumed octets.
func DecodeQoERVQoEReportingPaths(data []byte) (QoERVQoEReportingPaths, int, error) {
	var v QoERVQoEReportingPaths
	err := ngapcodec.DecodeParameter("QoERVQoEReportingPaths", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoERVQoEReportingPaths{}, 0, err
	}
	return v, len(data), err
}
