// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NGRANCGIExtIEsEntryPresent uint32

const (
	NGRANCGIExtIEsEntryPresentNothing NGRANCGIExtIEsEntryPresent = iota
	NGRANCGIExtIEsEntryPresentUnknown
)

type NGRANCGIExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      NGRANCGIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANCGIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-CGI-ExtIEsEntry", v)
}

// DecodeNGRANCGIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANCGIExtIEsEntry(data []byte) (NGRANCGIExtIEsEntry, int, error) {
	var v NGRANCGIExtIEsEntry
	err := ngapcodec.DecodeParameter("NGRAN-CGI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANCGIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
