// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MRBProgressInformationExtIEsEntryPresent uint32

const (
	MRBProgressInformationExtIEsEntryPresentNothing MRBProgressInformationExtIEsEntryPresent = iota
	MRBProgressInformationExtIEsEntryPresentUnknown
)

type MRBProgressInformationExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      MRBProgressInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MRBProgressInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MRB-ProgressInformation-ExtIEsEntry", v)
}

// DecodeMRBProgressInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMRBProgressInformationExtIEsEntry(data []byte) (MRBProgressInformationExtIEsEntry, int, error) {
	var v MRBProgressInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("MRB-ProgressInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MRBProgressInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
