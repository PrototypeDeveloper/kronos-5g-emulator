// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MDTConfigurationNRExtIEsEntryPresent uint32

const (
	MDTConfigurationNRExtIEsEntryPresentNothing MDTConfigurationNRExtIEsEntryPresent = iota
	MDTConfigurationNRExtIEsEntryPresentPNINPNAreaScopeofMDT
	MDTConfigurationNRExtIEsEntryPresentNetworkSliceAreaScopeofMDT
	MDTConfigurationNRExtIEsEntryPresentUnknown
)

type MDTConfigurationNRExtIEsEntry struct {
	ID                         ProtocolExtensionID
	Criticality                Criticality
	Present                    MDTConfigurationNRExtIEsEntryPresent
	PNINPNAreaScopeofMDT       PNINPNAreaScopeofMDT       `ngap:"id:409,criticality:ignore,presence:optional"`
	NetworkSliceAreaScopeofMDT NetworkSliceAreaScopeofMDT `ngap:"id:447,criticality:ignore,presence:optional"`
	UnknownValue               []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTConfigurationNRExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Configuration-NR-ExtIEsEntry", v)
}

// DecodeMDTConfigurationNRExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMDTConfigurationNRExtIEsEntry(data []byte) (MDTConfigurationNRExtIEsEntry, int, error) {
	var v MDTConfigurationNRExtIEsEntry
	err := ngapcodec.DecodeParameter("MDT-Configuration-NR-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTConfigurationNRExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
