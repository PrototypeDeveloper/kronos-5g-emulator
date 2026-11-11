// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AvailableBitrateReportThresholdItemExtIEsEntryPresent uint32

const (
	AvailableBitrateReportThresholdItemExtIEsEntryPresentNothing AvailableBitrateReportThresholdItemExtIEsEntryPresent = iota
	AvailableBitrateReportThresholdItemExtIEsEntryPresentUnknown
)

type AvailableBitrateReportThresholdItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AvailableBitrateReportThresholdItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AvailableBitrateReportThresholdItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AvailableBitrateReportThresholdItem-ExtIEsEntry", v)
}

// DecodeAvailableBitrateReportThresholdItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAvailableBitrateReportThresholdItemExtIEsEntry(data []byte) (AvailableBitrateReportThresholdItemExtIEsEntry, int, error) {
	var v AvailableBitrateReportThresholdItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AvailableBitrateReportThresholdItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AvailableBitrateReportThresholdItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
