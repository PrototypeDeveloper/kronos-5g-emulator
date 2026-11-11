// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LastVisitedPSCellInformationExtIEsEntryPresent uint32

const (
	LastVisitedPSCellInformationExtIEsEntryPresentNothing LastVisitedPSCellInformationExtIEsEntryPresent = iota
	LastVisitedPSCellInformationExtIEsEntryPresentSCGActivationTime
	LastVisitedPSCellInformationExtIEsEntryPresentUnknown
)

type LastVisitedPSCellInformationExtIEsEntry struct {
	ID                ProtocolExtensionID
	Criticality       Criticality
	Present           LastVisitedPSCellInformationExtIEsEntryPresent
	SCGActivationTime SCGActivationTime `ngap:"id:492,criticality:ignore,presence:optional"`
	UnknownValue      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedPSCellInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedPSCellInformation-ExtIEsEntry", v)
}

// DecodeLastVisitedPSCellInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedPSCellInformationExtIEsEntry(data []byte) (LastVisitedPSCellInformationExtIEsEntry, int, error) {
	var v LastVisitedPSCellInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("LastVisitedPSCellInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedPSCellInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
