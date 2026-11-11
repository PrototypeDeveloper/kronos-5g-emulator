// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type M6ConfigurationExtIEsEntryPresent uint32

const (
	M6ConfigurationExtIEsEntryPresentNothing M6ConfigurationExtIEsEntryPresent = iota
	M6ConfigurationExtIEsEntryPresentM6ReportAmount
	M6ConfigurationExtIEsEntryPresentExcessPacketDelayThresholdConfiguration
	M6ConfigurationExtIEsEntryPresentUnknown
)

type M6ConfigurationExtIEsEntry struct {
	ID                                      ProtocolExtensionID
	Criticality                             Criticality
	Present                                 M6ConfigurationExtIEsEntryPresent
	M6ReportAmount                          M6ReportAmountMDT                       `ngap:"id:338,criticality:ignore,presence:optional"`
	ExcessPacketDelayThresholdConfiguration ExcessPacketDelayThresholdConfiguration `ngap:"id:341,criticality:ignore,presence:optional"`
	UnknownValue                            []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M6ConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M6Configuration-ExtIEsEntry", v)
}

// DecodeM6ConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeM6ConfigurationExtIEsEntry(data []byte) (M6ConfigurationExtIEsEntry, int, error) {
	var v M6ConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("M6Configuration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return M6ConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
