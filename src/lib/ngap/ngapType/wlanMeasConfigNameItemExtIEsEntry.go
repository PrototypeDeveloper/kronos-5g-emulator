// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type WLANMeasConfigNameItemExtIEsEntryPresent uint32

const (
	WLANMeasConfigNameItemExtIEsEntryPresentNothing WLANMeasConfigNameItemExtIEsEntryPresent = iota
	WLANMeasConfigNameItemExtIEsEntryPresentUnknown
)

type WLANMeasConfigNameItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      WLANMeasConfigNameItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WLANMeasConfigNameItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WLANMeasConfigNameItem-ExtIEsEntry", v)
}

// DecodeWLANMeasConfigNameItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeWLANMeasConfigNameItemExtIEsEntry(data []byte) (WLANMeasConfigNameItemExtIEsEntry, int, error) {
	var v WLANMeasConfigNameItemExtIEsEntry
	err := ngapcodec.DecodeParameter("WLANMeasConfigNameItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANMeasConfigNameItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
