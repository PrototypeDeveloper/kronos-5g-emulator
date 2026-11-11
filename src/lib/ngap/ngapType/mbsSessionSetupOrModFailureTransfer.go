// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupOrModFailureTransfer represents ASN.1 MBSSessionSetupOrModFailureTransfer.
type MBSSessionSetupOrModFailureTransfer struct {
	Cause                     Cause
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []MBSSessionSetupOrModFailureTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupOrModFailureTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupOrModFailureTransfer", v)
}

// DecodeMBSSessionSetupOrModFailureTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupOrModFailureTransfer(data []byte) (MBSSessionSetupOrModFailureTransfer, int, error) {
	var v MBSSessionSetupOrModFailureTransfer
	err := ngapcodec.DecodeParameter("MBSSessionSetupOrModFailureTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupOrModFailureTransfer{}, 0, err
	}
	return v, len(data), err
}
