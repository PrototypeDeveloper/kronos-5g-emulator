// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LastVisitedCellInformationExtIEsEntryPresent uint32

const (
	LastVisitedCellInformationExtIEsEntryPresentNothing LastVisitedCellInformationExtIEsEntryPresent = iota
	LastVisitedCellInformationExtIEsEntryPresentUnknown
)

type LastVisitedCellInformationExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      LastVisitedCellInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedCellInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedCellInformation-ExtIEsEntry", v)
}

// DecodeLastVisitedCellInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedCellInformationExtIEsEntry(data []byte) (LastVisitedCellInformationExtIEsEntry, int, error) {
	var v LastVisitedCellInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("LastVisitedCellInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedCellInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
