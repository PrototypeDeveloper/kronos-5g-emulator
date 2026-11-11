// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MDTConfigurationExtIEsEntryPresent uint32

const (
	MDTConfigurationExtIEsEntryPresentNothing MDTConfigurationExtIEsEntryPresent = iota
	MDTConfigurationExtIEsEntryPresentMNOnlyMDTCollection
	MDTConfigurationExtIEsEntryPresentUnknown
)

type MDTConfigurationExtIEsEntry struct {
	ID                  ProtocolExtensionID
	Criticality         Criticality
	Present             MDTConfigurationExtIEsEntryPresent
	MNOnlyMDTCollection MNOnlyMDTCollection `ngap:"id:433,criticality:ignore,presence:optional"`
	UnknownValue        []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Configuration-ExtIEsEntry", v)
}

// DecodeMDTConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMDTConfigurationExtIEsEntry(data []byte) (MDTConfigurationExtIEsEntry, int, error) {
	var v MDTConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("MDT-Configuration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
