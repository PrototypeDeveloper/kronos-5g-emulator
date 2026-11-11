// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSServiceAreaInformationExtIEsEntryPresent uint32

const (
	MBSServiceAreaInformationExtIEsEntryPresentNothing MBSServiceAreaInformationExtIEsEntryPresent = iota
	MBSServiceAreaInformationExtIEsEntryPresentMBSIntendedServiceAreaList
	MBSServiceAreaInformationExtIEsEntryPresentUnknown
)

type MBSServiceAreaInformationExtIEsEntry struct {
	ID                         ProtocolExtensionID
	Criticality                Criticality
	Present                    MBSServiceAreaInformationExtIEsEntryPresent
	MBSIntendedServiceAreaList MBSIntendedServiceAreaList `ngap:"id:449,criticality:ignore,presence:optional"`
	UnknownValue               []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceAreaInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceAreaInformation-ExtIEsEntry", v)
}

// DecodeMBSServiceAreaInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceAreaInformationExtIEsEntry(data []byte) (MBSServiceAreaInformationExtIEsEntry, int, error) {
	var v MBSServiceAreaInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-ServiceAreaInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceAreaInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
