// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EventL1LoggedMDTConfigExtIEsEntryPresent uint32

const (
	EventL1LoggedMDTConfigExtIEsEntryPresentNothing EventL1LoggedMDTConfigExtIEsEntryPresent = iota
	EventL1LoggedMDTConfigExtIEsEntryPresentUnknown
)

type EventL1LoggedMDTConfigExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EventL1LoggedMDTConfigExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EventL1LoggedMDTConfigExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EventL1LoggedMDTConfig-ExtIEsEntry", v)
}

// DecodeEventL1LoggedMDTConfigExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEventL1LoggedMDTConfigExtIEsEntry(data []byte) (EventL1LoggedMDTConfigExtIEsEntry, int, error) {
	var v EventL1LoggedMDTConfigExtIEsEntry
	err := ngapcodec.DecodeParameter("EventL1LoggedMDTConfig-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EventL1LoggedMDTConfigExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
