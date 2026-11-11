// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBStatusUL12 represents ASN.1 DRBStatusUL12.
type DRBStatusUL12 struct {
	ULCOUNTValue                 COUNTValueForPDCPSN12
	HasReceiveStatusOfULPDCPSDUs bool
	ReceiveStatusOfULPDCPSDUs    BitString `aper:"sizeLB:1,sizeUB:2048"`
	HasIEExtension               bool
	IEExtension                  []DRBStatusUL12ExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusUL12) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DRBStatusUL12", v) }

// DecodeDRBStatusUL12 decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusUL12(data []byte) (DRBStatusUL12, int, error) {
	var v DRBStatusUL12
	err := ngapcodec.DecodeParameter("DRBStatusUL12", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusUL12{}, 0, err
	}
	return v, len(data), err
}
