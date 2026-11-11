// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SliceSupportListQMC represents ASN.1 SliceSupportListQMC.
type SliceSupportListQMC struct {
	List []SliceSupportQMCItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SliceSupportListQMC) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SliceSupportListQMC", v)
}

// DecodeSliceSupportListQMC decodes a standalone APER value and reports the consumed octets.
func DecodeSliceSupportListQMC(data []byte) (SliceSupportListQMC, int, error) {
	var v SliceSupportListQMC
	err := ngapcodec.DecodeParameter("SliceSupportListQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return SliceSupportListQMC{}, 0, err
	}
	return v, len(data), err
}
