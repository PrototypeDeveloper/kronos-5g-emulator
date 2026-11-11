// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBStatusDL12 represents ASN.1 DRBStatusDL12.
type DRBStatusDL12 struct {
	DLCOUNTValue      COUNTValueForPDCPSN12
	HasIEExtension    bool
	IEExtension       []DRBStatusDL12ExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusDL12) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DRBStatusDL12", v) }

// DecodeDRBStatusDL12 decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusDL12(data []byte) (DRBStatusDL12, int, error) {
	var v DRBStatusDL12
	err := ngapcodec.DecodeParameter("DRBStatusDL12", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusDL12{}, 0, err
	}
	return v, len(data), err
}
