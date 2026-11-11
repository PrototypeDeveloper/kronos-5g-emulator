// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EUTRACGIExtIEsEntryPresent uint32

const (
	EUTRACGIExtIEsEntryPresentNothing EUTRACGIExtIEsEntryPresent = iota
	EUTRACGIExtIEsEntryPresentUnknown
)

type EUTRACGIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EUTRACGIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRACGIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRA-CGI-ExtIEsEntry", v)
}

// DecodeEUTRACGIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRACGIExtIEsEntry(data []byte) (EUTRACGIExtIEsEntry, int, error) {
	var v EUTRACGIExtIEsEntry
	err := ngapcodec.DecodeParameter("EUTRA-CGI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRACGIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
