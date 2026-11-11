// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNTAIListforMDT represents ASN.1 SNPNTAIListforMDT.
type SNPNTAIListforMDT struct {
	List []SNPNTAIListforMDTItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNTAIListforMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNTAIListforMDT", v)
}

// DecodeSNPNTAIListforMDT decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNTAIListforMDT(data []byte) (SNPNTAIListforMDT, int, error) {
	var v SNPNTAIListforMDT
	err := ngapcodec.DecodeParameter("SNPNTAIListforMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNTAIListforMDT{}, 0, err
	}
	return v, len(data), err
}
