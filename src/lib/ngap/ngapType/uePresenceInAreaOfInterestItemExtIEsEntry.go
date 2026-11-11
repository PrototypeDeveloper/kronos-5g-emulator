// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEPresenceInAreaOfInterestItemExtIEsEntryPresent uint32

const (
	UEPresenceInAreaOfInterestItemExtIEsEntryPresentNothing UEPresenceInAreaOfInterestItemExtIEsEntryPresent = iota
	UEPresenceInAreaOfInterestItemExtIEsEntryPresentUnknown
)

type UEPresenceInAreaOfInterestItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEPresenceInAreaOfInterestItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPresenceInAreaOfInterestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEPresenceInAreaOfInterestItem-ExtIEsEntry", v)
}

// DecodeUEPresenceInAreaOfInterestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEPresenceInAreaOfInterestItemExtIEsEntry(data []byte) (UEPresenceInAreaOfInterestItemExtIEsEntry, int, error) {
	var v UEPresenceInAreaOfInterestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UEPresenceInAreaOfInterestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPresenceInAreaOfInterestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
