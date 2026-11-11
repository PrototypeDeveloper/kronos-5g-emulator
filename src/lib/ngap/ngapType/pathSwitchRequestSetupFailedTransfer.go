// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PathSwitchRequestSetupFailedTransfer represents ASN.1 PathSwitchRequestSetupFailedTransfer.
type PathSwitchRequestSetupFailedTransfer struct {
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []PathSwitchRequestSetupFailedTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PathSwitchRequestSetupFailedTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PathSwitchRequestSetupFailedTransfer", v)
}

// DecodePathSwitchRequestSetupFailedTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePathSwitchRequestSetupFailedTransfer(data []byte) (PathSwitchRequestSetupFailedTransfer, int, error) {
	var v PathSwitchRequestSetupFailedTransfer
	err := ngapcodec.DecodeParameter("PathSwitchRequestSetupFailedTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PathSwitchRequestSetupFailedTransfer{}, 0, err
	}
	return v, len(data), err
}
