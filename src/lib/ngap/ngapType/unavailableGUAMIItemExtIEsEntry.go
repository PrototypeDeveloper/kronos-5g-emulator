// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UnavailableGUAMIItemExtIEsEntryPresent uint32

const (
	UnavailableGUAMIItemExtIEsEntryPresentNothing UnavailableGUAMIItemExtIEsEntryPresent = iota
	UnavailableGUAMIItemExtIEsEntryPresentExtendedBackupAMFName
	UnavailableGUAMIItemExtIEsEntryPresentUnknown
)

type UnavailableGUAMIItemExtIEsEntry struct {
	ID                    ProtocolExtensionID
	Criticality           Criticality
	Present               UnavailableGUAMIItemExtIEsEntryPresent
	ExtendedBackupAMFName ExtendedAMFName `ngap:"id:442,criticality:ignore,presence:optional"`
	UnknownValue          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UnavailableGUAMIItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UnavailableGUAMIItem-ExtIEsEntry", v)
}

// DecodeUnavailableGUAMIItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUnavailableGUAMIItemExtIEsEntry(data []byte) (UnavailableGUAMIItemExtIEsEntry, int, error) {
	var v UnavailableGUAMIItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UnavailableGUAMIItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UnavailableGUAMIItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
