// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SONInformationExtIEsEntryPresent uint32

const (
	SONInformationExtIEsEntryPresentNothing SONInformationExtIEsEntryPresent = iota
	SONInformationExtIEsEntryPresentSONInformationReport
	SONInformationExtIEsEntryPresentUnknown
)

type SONInformationExtIEsEntry struct {
	ID                   ProtocolIEID
	Criticality          Criticality
	Present              SONInformationExtIEsEntryPresent
	SONInformationReport SONInformationReport `ngap:"id:252,criticality:ignore,presence:mandatory"`
	UnknownValue         []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONInformation-ExtIEsEntry", v)
}

// DecodeSONInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSONInformationExtIEsEntry(data []byte) (SONInformationExtIEsEntry, int, error) {
	var v SONInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("SONInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
