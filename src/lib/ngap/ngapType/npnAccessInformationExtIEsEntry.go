// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NPNAccessInformationExtIEsEntryPresent uint32

const (
	NPNAccessInformationExtIEsEntryPresentNothing NPNAccessInformationExtIEsEntryPresent = iota
	NPNAccessInformationExtIEsEntryPresentUnknown
)

type NPNAccessInformationExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      NPNAccessInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NPNAccessInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NPN-AccessInformation-ExtIEsEntry", v)
}

// DecodeNPNAccessInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNPNAccessInformationExtIEsEntry(data []byte) (NPNAccessInformationExtIEsEntry, int, error) {
	var v NPNAccessInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("NPN-AccessInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NPNAccessInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
