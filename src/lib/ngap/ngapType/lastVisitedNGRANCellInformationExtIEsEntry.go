// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LastVisitedNGRANCellInformationExtIEsEntryPresent uint32

const (
	LastVisitedNGRANCellInformationExtIEsEntryPresentNothing LastVisitedNGRANCellInformationExtIEsEntryPresent = iota
	LastVisitedNGRANCellInformationExtIEsEntryPresentLastVisitedPSCellList
	LastVisitedNGRANCellInformationExtIEsEntryPresentUnknown
)

type LastVisitedNGRANCellInformationExtIEsEntry struct {
	ID                    ProtocolExtensionID
	Criticality           Criticality
	Present               LastVisitedNGRANCellInformationExtIEsEntryPresent
	LastVisitedPSCellList LastVisitedPSCellList `ngap:"id:289,criticality:ignore,presence:optional"`
	UnknownValue          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedNGRANCellInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedNGRANCellInformation-ExtIEsEntry", v)
}

// DecodeLastVisitedNGRANCellInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedNGRANCellInformationExtIEsEntry(data []byte) (LastVisitedNGRANCellInformationExtIEsEntry, int, error) {
	var v LastVisitedNGRANCellInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("LastVisitedNGRANCellInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedNGRANCellInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
