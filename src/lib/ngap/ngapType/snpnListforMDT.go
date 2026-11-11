// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNListforMDT represents ASN.1 SNPNListforMDT.
type SNPNListforMDT struct {
	List []SNPNListforMDTItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNListforMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNListforMDT", v)
}

// DecodeSNPNListforMDT decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNListforMDT(data []byte) (SNPNListforMDT, int, error) {
	var v SNPNListforMDT
	err := ngapcodec.DecodeParameter("SNPNListforMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNListforMDT{}, 0, err
	}
	return v, len(data), err
}
