// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SONConfigurationTransfer represents ASN.1 SONConfigurationTransfer.
type SONConfigurationTransfer struct {
	TargetRANNodeIDSON        TargetRANNodeIDSON
	SourceRANNodeID           SourceRANNodeID
	SONInformation            SONInformation
	HasXnTNLConfigurationInfo bool
	XnTNLConfigurationInfo    XnTNLConfigurationInfo
	HasIEExtensions           bool
	IEExtensions              []SONConfigurationTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONConfigurationTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONConfigurationTransfer", v)
}

// DecodeSONConfigurationTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeSONConfigurationTransfer(data []byte) (SONConfigurationTransfer, int, error) {
	var v SONConfigurationTransfer
	err := ngapcodec.DecodeParameter("SONConfigurationTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONConfigurationTransfer{}, 0, err
	}
	return v, len(data), err
}
