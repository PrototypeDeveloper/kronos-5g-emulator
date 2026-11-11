// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type M1ConfigurationExtIEsEntryPresent uint32

const (
	M1ConfigurationExtIEsEntryPresentNothing M1ConfigurationExtIEsEntryPresent = iota
	M1ConfigurationExtIEsEntryPresentIncludeBeamMeasurementsIndication
	M1ConfigurationExtIEsEntryPresentBeamMeasurementsReportConfiguration
	M1ConfigurationExtIEsEntryPresentUnknown
)

type M1ConfigurationExtIEsEntry struct {
	ID                                  ProtocolExtensionID
	Criticality                         Criticality
	Present                             M1ConfigurationExtIEsEntryPresent
	IncludeBeamMeasurementsIndication   IncludeBeamMeasurementsIndication   `ngap:"id:340,criticality:ignore,presence:optional"`
	BeamMeasurementsReportConfiguration BeamMeasurementsReportConfiguration `ngap:"id:361,criticality:ignore,presence:conditional"`
	UnknownValue                        []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1ConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1Configuration-ExtIEsEntry", v)
}

// DecodeM1ConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeM1ConfigurationExtIEsEntry(data []byte) (M1ConfigurationExtIEsEntry, int, error) {
	var v M1ConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("M1Configuration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1ConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
