// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CommandFailureTransfer represents ASN.1 CommandFailureTransfer.
type CommandFailureTransfer struct {
	CorrelationIdentifier     AIoTCorrelationIdentifier
	RANAIOTDeviceNGAPID       RANAIOTDeviceNGAPID
	Cause                     Cause
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []CommandFailureTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandFailureTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandFailureTransfer", v)
}

// DecodeCommandFailureTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeCommandFailureTransfer(data []byte) (CommandFailureTransfer, int, error) {
	var v CommandFailureTransfer
	err := ngapcodec.DecodeParameter("CommandFailureTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandFailureTransfer{}, 0, err
	}
	return v, len(data), err
}
