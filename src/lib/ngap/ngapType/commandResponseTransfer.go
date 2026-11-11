// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CommandResponseTransfer represents ASN.1 CommandResponseTransfer.
type CommandResponseTransfer struct {
	CorrelationIdentifier     AIoTCorrelationIdentifier
	RANAIOTDeviceNGAPID       RANAIOTDeviceNGAPID
	AIoTNASPDU                AIoTNASPDU
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    CriticalityDiagnostics
	HasIEExtensions           bool
	IEExtensions              []CommandResponseTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandResponseTransfer", v)
}

// DecodeCommandResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeCommandResponseTransfer(data []byte) (CommandResponseTransfer, int, error) {
	var v CommandResponseTransfer
	err := ngapcodec.DecodeParameter("CommandResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
