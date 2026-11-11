// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSDistributionSetupUnsuccessfulTransferExtIEsEntryPresent uint32

const (
	MBSDistributionSetupUnsuccessfulTransferExtIEsEntryPresentNothing MBSDistributionSetupUnsuccessfulTransferExtIEsEntryPresent = iota
	MBSDistributionSetupUnsuccessfulTransferExtIEsEntryPresentUnknown
)

type MBSDistributionSetupUnsuccessfulTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSDistributionSetupUnsuccessfulTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupUnsuccessfulTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupUnsuccessfulTransfer-ExtIEsEntry", v)
}

// DecodeMBSDistributionSetupUnsuccessfulTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupUnsuccessfulTransferExtIEsEntry(data []byte) (MBSDistributionSetupUnsuccessfulTransferExtIEsEntry, int, error) {
	var v MBSDistributionSetupUnsuccessfulTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupUnsuccessfulTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupUnsuccessfulTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
