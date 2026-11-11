// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SourceToTargetAMFInformationRerouteExtIEsEntryPresent uint32

const (
	SourceToTargetAMFInformationRerouteExtIEsEntryPresentNothing SourceToTargetAMFInformationRerouteExtIEsEntryPresent = iota
	SourceToTargetAMFInformationRerouteExtIEsEntryPresentUnknown
)

type SourceToTargetAMFInformationRerouteExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SourceToTargetAMFInformationRerouteExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceToTargetAMFInformationRerouteExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceToTarget-AMFInformationReroute-ExtIEsEntry", v)
}

// DecodeSourceToTargetAMFInformationRerouteExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSourceToTargetAMFInformationRerouteExtIEsEntry(data []byte) (SourceToTargetAMFInformationRerouteExtIEsEntry, int, error) {
	var v SourceToTargetAMFInformationRerouteExtIEsEntry
	err := ngapcodec.DecodeParameter("SourceToTarget-AMFInformationReroute-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceToTargetAMFInformationRerouteExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
