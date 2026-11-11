// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBStatusDL18 represents ASN.1 DRBStatusDL18.
type DRBStatusDL18 struct {
	DLCOUNTValue      COUNTValueForPDCPSN18
	HasIEExtension    bool
	IEExtension       []DRBStatusDL18ExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusDL18) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DRBStatusDL18", v) }

// DecodeDRBStatusDL18 decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusDL18(data []byte) (DRBStatusDL18, int, error) {
	var v DRBStatusDL18
	err := ngapcodec.DecodeParameter("DRBStatusDL18", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusDL18{}, 0, err
	}
	return v, len(data), err
}
