// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QoERVQoEReportingPathsExtIEsEntryPresent uint32

const (
	QoERVQoEReportingPathsExtIEsEntryPresentNothing QoERVQoEReportingPathsExtIEsEntryPresent = iota
	QoERVQoEReportingPathsExtIEsEntryPresentUnknown
)

type QoERVQoEReportingPathsExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QoERVQoEReportingPathsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoERVQoEReportingPathsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoERVQoEReportingPaths-ExtIEsEntry", v)
}

// DecodeQoERVQoEReportingPathsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQoERVQoEReportingPathsExtIEsEntry(data []byte) (QoERVQoEReportingPathsExtIEsEntry, int, error) {
	var v QoERVQoEReportingPathsExtIEsEntry
	err := ngapcodec.DecodeParameter("QoERVQoEReportingPaths-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoERVQoEReportingPathsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
