// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EUTRANRadioResourceStatusExtIEsEntryPresent uint32

const (
	EUTRANRadioResourceStatusExtIEsEntryPresentNothing EUTRANRadioResourceStatusExtIEsEntryPresent = iota
	EUTRANRadioResourceStatusExtIEsEntryPresentUnknown
)

type EUTRANRadioResourceStatusExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EUTRANRadioResourceStatusExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANRadioResourceStatusExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-RadioResourceStatus-ExtIEsEntry", v)
}

// DecodeEUTRANRadioResourceStatusExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANRadioResourceStatusExtIEsEntry(data []byte) (EUTRANRadioResourceStatusExtIEsEntry, int, error) {
	var v EUTRANRadioResourceStatusExtIEsEntry
	err := ngapcodec.DecodeParameter("EUTRAN-RadioResourceStatus-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANRadioResourceStatusExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
