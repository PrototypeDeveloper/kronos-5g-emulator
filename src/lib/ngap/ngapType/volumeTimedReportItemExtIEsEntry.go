// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type VolumeTimedReportItemExtIEsEntryPresent uint32

const (
	VolumeTimedReportItemExtIEsEntryPresentNothing VolumeTimedReportItemExtIEsEntryPresent = iota
	VolumeTimedReportItemExtIEsEntryPresentUnknown
)

type VolumeTimedReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      VolumeTimedReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v VolumeTimedReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("VolumeTimedReport-Item-ExtIEsEntry", v)
}

// DecodeVolumeTimedReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeVolumeTimedReportItemExtIEsEntry(data []byte) (VolumeTimedReportItemExtIEsEntry, int, error) {
	var v VolumeTimedReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("VolumeTimedReport-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return VolumeTimedReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
