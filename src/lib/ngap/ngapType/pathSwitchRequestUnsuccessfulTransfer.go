// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PathSwitchRequestUnsuccessfulTransfer represents ASN.1 PathSwitchRequestUnsuccessfulTransfer.
type PathSwitchRequestUnsuccessfulTransfer struct {
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []PathSwitchRequestUnsuccessfulTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PathSwitchRequestUnsuccessfulTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PathSwitchRequestUnsuccessfulTransfer", v)
}

// DecodePathSwitchRequestUnsuccessfulTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePathSwitchRequestUnsuccessfulTransfer(data []byte) (PathSwitchRequestUnsuccessfulTransfer, int, error) {
	var v PathSwitchRequestUnsuccessfulTransfer
	err := ngapcodec.DecodeParameter("PathSwitchRequestUnsuccessfulTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PathSwitchRequestUnsuccessfulTransfer{}, 0, err
	}
	return v, len(data), err
}
