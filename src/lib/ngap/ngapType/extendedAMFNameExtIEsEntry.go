// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ExtendedAMFNameExtIEsEntryPresent uint32

const (
	ExtendedAMFNameExtIEsEntryPresentNothing ExtendedAMFNameExtIEsEntryPresent = iota
	ExtendedAMFNameExtIEsEntryPresentUnknown
)

type ExtendedAMFNameExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ExtendedAMFNameExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedAMFNameExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Extended-AMFName-ExtIEsEntry", v)
}

// DecodeExtendedAMFNameExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedAMFNameExtIEsEntry(data []byte) (ExtendedAMFNameExtIEsEntry, int, error) {
	var v ExtendedAMFNameExtIEsEntry
	err := ngapcodec.DecodeParameter("Extended-AMFName-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedAMFNameExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
