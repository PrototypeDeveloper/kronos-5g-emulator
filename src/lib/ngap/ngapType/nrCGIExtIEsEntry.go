// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRCGIExtIEsEntryPresent uint32

const (
	NRCGIExtIEsEntryPresentNothing NRCGIExtIEsEntryPresent = iota
	NRCGIExtIEsEntryPresentUnknown
)

type NRCGIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRCGIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRCGIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-CGI-ExtIEsEntry", v)
}

// DecodeNRCGIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRCGIExtIEsEntry(data []byte) (NRCGIExtIEsEntry, int, error) {
	var v NRCGIExtIEsEntry
	err := ngapcodec.DecodeParameter("NR-CGI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRCGIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
