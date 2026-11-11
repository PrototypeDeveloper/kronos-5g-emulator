// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ENDCSONConfigurationTransfer represents ASN.1 EN-DCSONConfigurationTransfer.
type ENDCSONConfigurationTransfer struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ENDCSONConfigurationTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EN-DCSONConfigurationTransfer", v)
}

// DecodeENDCSONConfigurationTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeENDCSONConfigurationTransfer(data []byte) (ENDCSONConfigurationTransfer, int, error) {
	var v ENDCSONConfigurationTransfer
	err := ngapcodec.DecodeParameter("EN-DCSONConfigurationTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return ENDCSONConfigurationTransfer{}, 0, err
	}
	return v, len(data), err
}
