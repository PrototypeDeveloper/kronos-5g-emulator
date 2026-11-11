// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSDistributionReleaseRequesTransferExtIEsEntryPresent uint32

const (
	MBSDistributionReleaseRequesTransferExtIEsEntryPresentNothing MBSDistributionReleaseRequesTransferExtIEsEntryPresent = iota
	MBSDistributionReleaseRequesTransferExtIEsEntryPresentUnknown
)

type MBSDistributionReleaseRequesTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSDistributionReleaseRequesTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionReleaseRequesTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionReleaseRequesTransfer-ExtIEsEntry", v)
}

// DecodeMBSDistributionReleaseRequesTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionReleaseRequesTransferExtIEsEntry(data []byte) (MBSDistributionReleaseRequesTransferExtIEsEntry, int, error) {
	var v MBSDistributionReleaseRequesTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-DistributionReleaseRequesTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionReleaseRequesTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
