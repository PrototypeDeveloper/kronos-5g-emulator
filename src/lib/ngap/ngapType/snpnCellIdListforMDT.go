// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNCellIdListforMDT represents ASN.1 SNPNCellIdListforMDT.
type SNPNCellIdListforMDT struct {
	List []SNPNCellIdListforMDTItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNCellIdListforMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNCellIdListforMDT", v)
}

// DecodeSNPNCellIdListforMDT decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNCellIdListforMDT(data []byte) (SNPNCellIdListforMDT, int, error) {
	var v SNPNCellIdListforMDT
	err := ngapcodec.DecodeParameter("SNPNCellIdListforMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNCellIdListforMDT{}, 0, err
	}
	return v, len(data), err
}
