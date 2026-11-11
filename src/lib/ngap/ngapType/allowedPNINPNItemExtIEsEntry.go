// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AllowedPNINPNItemExtIEsEntryPresent uint32

const (
	AllowedPNINPNItemExtIEsEntryPresentNothing AllowedPNINPNItemExtIEsEntryPresent = iota
	AllowedPNINPNItemExtIEsEntryPresentUnknown
)

type AllowedPNINPNItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AllowedPNINPNItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedPNINPNItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Allowed-PNI-NPN-Item-ExtIEsEntry", v)
}

// DecodeAllowedPNINPNItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedPNINPNItemExtIEsEntry(data []byte) (AllowedPNINPNItemExtIEsEntry, int, error) {
	var v AllowedPNINPNItemExtIEsEntry
	err := ngapcodec.DecodeParameter("Allowed-PNI-NPN-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedPNINPNItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
