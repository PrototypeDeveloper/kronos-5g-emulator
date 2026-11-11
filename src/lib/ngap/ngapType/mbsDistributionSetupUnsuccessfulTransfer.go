// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDistributionSetupUnsuccessfulTransfer represents ASN.1 MBS-DistributionSetupUnsuccessfulTransfer.
type MBSDistributionSetupUnsuccessfulTransfer struct {
	MBSSessionID              MBSSessionID
	HasMBSAreaSessionID       bool
	MBSAreaSessionID          MBSAreaSessionID
	Cause                     Cause
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []MBSDistributionSetupUnsuccessfulTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupUnsuccessfulTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupUnsuccessfulTransfer", v)
}

// DecodeMBSDistributionSetupUnsuccessfulTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupUnsuccessfulTransfer(data []byte) (MBSDistributionSetupUnsuccessfulTransfer, int, error) {
	var v MBSDistributionSetupUnsuccessfulTransfer
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupUnsuccessfulTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupUnsuccessfulTransfer{}, 0, err
	}
	return v, len(data), err
}
