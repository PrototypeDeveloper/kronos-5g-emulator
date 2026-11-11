// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemSONConfigurationTransfer represents ASN.1 IntersystemSONConfigurationTransfer.
type IntersystemSONConfigurationTransfer struct {
	TransferType              IntersystemSONTransferType
	IntersystemSONInformation IntersystemSONInformation
	HasIEExtensions           bool
	IEExtensions              []IntersystemSONConfigurationTransferExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONConfigurationTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONConfigurationTransfer", v)
}

// DecodeIntersystemSONConfigurationTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONConfigurationTransfer(data []byte) (IntersystemSONConfigurationTransfer, int, error) {
	var v IntersystemSONConfigurationTransfer
	err := ngapcodec.DecodeParameter("IntersystemSONConfigurationTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONConfigurationTransfer{}, 0, err
	}
	return v, len(data), err
}
