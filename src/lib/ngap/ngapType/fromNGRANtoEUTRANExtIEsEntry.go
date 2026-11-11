// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FromNGRANtoEUTRANExtIEsEntryPresent uint32

const (
	FromNGRANtoEUTRANExtIEsEntryPresentNothing FromNGRANtoEUTRANExtIEsEntryPresent = iota
	FromNGRANtoEUTRANExtIEsEntryPresentUnknown
)

type FromNGRANtoEUTRANExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      FromNGRANtoEUTRANExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FromNGRANtoEUTRANExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FromNGRANtoEUTRAN-ExtIEsEntry", v)
}

// DecodeFromNGRANtoEUTRANExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFromNGRANtoEUTRANExtIEsEntry(data []byte) (FromNGRANtoEUTRANExtIEsEntry, int, error) {
	var v FromNGRANtoEUTRANExtIEsEntry
	err := ngapcodec.DecodeParameter("FromNGRANtoEUTRAN-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FromNGRANtoEUTRANExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
