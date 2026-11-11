// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSDistributionSetupResponseTransferExtIEsEntryPresent uint32

const (
	MBSDistributionSetupResponseTransferExtIEsEntryPresentNothing MBSDistributionSetupResponseTransferExtIEsEntryPresent = iota
	MBSDistributionSetupResponseTransferExtIEsEntryPresentUnknown
)

type MBSDistributionSetupResponseTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSDistributionSetupResponseTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupResponseTransfer-ExtIEsEntry", v)
}

// DecodeMBSDistributionSetupResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupResponseTransferExtIEsEntry(data []byte) (MBSDistributionSetupResponseTransferExtIEsEntry, int, error) {
	var v MBSDistributionSetupResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
