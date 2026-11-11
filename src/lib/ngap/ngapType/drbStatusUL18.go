// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBStatusUL18 represents ASN.1 DRBStatusUL18.
type DRBStatusUL18 struct {
	ULCOUNTValue                 COUNTValueForPDCPSN18
	HasReceiveStatusOfULPDCPSDUs bool
	ReceiveStatusOfULPDCPSDUs    BitString `aper:"sizeLB:1,sizeUB:131072"`
	HasIEExtension               bool
	IEExtension                  []DRBStatusUL18ExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusUL18) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DRBStatusUL18", v) }

// DecodeDRBStatusUL18 decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusUL18(data []byte) (DRBStatusUL18, int, error) {
	var v DRBStatusUL18
	err := ngapcodec.DecodeParameter("DRBStatusUL18", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusUL18{}, 0, err
	}
	return v, len(data), err
}
