// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ServedGUAMIItemExtIEsEntryPresent uint32

const (
	ServedGUAMIItemExtIEsEntryPresentNothing ServedGUAMIItemExtIEsEntryPresent = iota
	ServedGUAMIItemExtIEsEntryPresentGUAMIType
	ServedGUAMIItemExtIEsEntryPresentExtendedBackupAMFName
	ServedGUAMIItemExtIEsEntryPresentUnknown
)

type ServedGUAMIItemExtIEsEntry struct {
	ID                    ProtocolExtensionID
	Criticality           Criticality
	Present               ServedGUAMIItemExtIEsEntryPresent
	GUAMIType             GUAMIType       `ngap:"id:176,criticality:ignore,presence:optional"`
	ExtendedBackupAMFName ExtendedAMFName `ngap:"id:442,criticality:ignore,presence:optional"`
	UnknownValue          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ServedGUAMIItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ServedGUAMIItem-ExtIEsEntry", v)
}

// DecodeServedGUAMIItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeServedGUAMIItemExtIEsEntry(data []byte) (ServedGUAMIItemExtIEsEntry, int, error) {
	var v ServedGUAMIItemExtIEsEntry
	err := ngapcodec.DecodeParameter("ServedGUAMIItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ServedGUAMIItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
