// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ResourceStatusReportingSystemExtIEsEntryPresent uint32

const (
	ResourceStatusReportingSystemExtIEsEntryPresentNothing ResourceStatusReportingSystemExtIEsEntryPresent = iota
	ResourceStatusReportingSystemExtIEsEntryPresentUnknown
)

type ResourceStatusReportingSystemExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ResourceStatusReportingSystemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ResourceStatusReportingSystemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ResourceStatusReportingSystem-ExtIEsEntry", v)
}

// DecodeResourceStatusReportingSystemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeResourceStatusReportingSystemExtIEsEntry(data []byte) (ResourceStatusReportingSystemExtIEsEntry, int, error) {
	var v ResourceStatusReportingSystemExtIEsEntry
	err := ngapcodec.DecodeParameter("ResourceStatusReportingSystem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ResourceStatusReportingSystemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
