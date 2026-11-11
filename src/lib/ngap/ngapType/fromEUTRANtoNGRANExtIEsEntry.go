// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FromEUTRANtoNGRANExtIEsEntryPresent uint32

const (
	FromEUTRANtoNGRANExtIEsEntryPresentNothing FromEUTRANtoNGRANExtIEsEntryPresent = iota
	FromEUTRANtoNGRANExtIEsEntryPresentUnknown
)

type FromEUTRANtoNGRANExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      FromEUTRANtoNGRANExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FromEUTRANtoNGRANExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FromEUTRANtoNGRAN-ExtIEsEntry", v)
}

// DecodeFromEUTRANtoNGRANExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFromEUTRANtoNGRANExtIEsEntry(data []byte) (FromEUTRANtoNGRANExtIEsEntry, int, error) {
	var v FromEUTRANtoNGRANExtIEsEntry
	err := ngapcodec.DecodeParameter("FromEUTRANtoNGRAN-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FromEUTRANtoNGRANExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
