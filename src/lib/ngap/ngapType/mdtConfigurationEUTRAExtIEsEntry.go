// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MDTConfigurationEUTRAExtIEsEntryPresent uint32

const (
	MDTConfigurationEUTRAExtIEsEntryPresentNothing MDTConfigurationEUTRAExtIEsEntryPresent = iota
	MDTConfigurationEUTRAExtIEsEntryPresentUnknown
)

type MDTConfigurationEUTRAExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MDTConfigurationEUTRAExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTConfigurationEUTRAExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Configuration-EUTRA-ExtIEsEntry", v)
}

// DecodeMDTConfigurationEUTRAExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMDTConfigurationEUTRAExtIEsEntry(data []byte) (MDTConfigurationEUTRAExtIEsEntry, int, error) {
	var v MDTConfigurationEUTRAExtIEsEntry
	err := ngapcodec.DecodeParameter("MDT-Configuration-EUTRA-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTConfigurationEUTRAExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
