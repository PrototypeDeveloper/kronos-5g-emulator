// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ExtendedRANNodeNameExtIEsEntryPresent uint32

const (
	ExtendedRANNodeNameExtIEsEntryPresentNothing ExtendedRANNodeNameExtIEsEntryPresent = iota
	ExtendedRANNodeNameExtIEsEntryPresentUnknown
)

type ExtendedRANNodeNameExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ExtendedRANNodeNameExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedRANNodeNameExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Extended-RANNodeName-ExtIEsEntry", v)
}

// DecodeExtendedRANNodeNameExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedRANNodeNameExtIEsEntry(data []byte) (ExtendedRANNodeNameExtIEsEntry, int, error) {
	var v ExtendedRANNodeNameExtIEsEntry
	err := ngapcodec.DecodeParameter("Extended-RANNodeName-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedRANNodeNameExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
