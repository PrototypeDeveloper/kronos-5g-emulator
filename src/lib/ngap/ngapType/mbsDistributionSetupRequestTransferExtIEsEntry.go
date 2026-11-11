// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSDistributionSetupRequestTransferExtIEsEntryPresent uint32

const (
	MBSDistributionSetupRequestTransferExtIEsEntryPresentNothing MBSDistributionSetupRequestTransferExtIEsEntryPresent = iota
	MBSDistributionSetupRequestTransferExtIEsEntryPresentTAIMBSSupportList
	MBSDistributionSetupRequestTransferExtIEsEntryPresentUnknown
)

type MBSDistributionSetupRequestTransferExtIEsEntry struct {
	ID                ProtocolExtensionID
	Criticality       Criticality
	Present           MBSDistributionSetupRequestTransferExtIEsEntryPresent
	TAIMBSSupportList TAIMBSSupportList `ngap:"id:441,criticality:ignore,presence:optional"`
	UnknownValue      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupRequestTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupRequestTransfer-ExtIEsEntry", v)
}

// DecodeMBSDistributionSetupRequestTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupRequestTransferExtIEsEntry(data []byte) (MBSDistributionSetupRequestTransferExtIEsEntry, int, error) {
	var v MBSDistributionSetupRequestTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupRequestTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupRequestTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
