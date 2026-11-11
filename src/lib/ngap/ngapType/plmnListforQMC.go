// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PLMNListforQMC represents ASN.1 PLMNListforQMC.
type PLMNListforQMC struct {
	List []PLMNIdentity `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PLMNListforQMC) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PLMNListforQMC", v)
}

// DecodePLMNListforQMC decodes a standalone APER value and reports the consumed octets.
func DecodePLMNListforQMC(data []byte) (PLMNListforQMC, int, error) {
	var v PLMNListforQMC
	err := ngapcodec.DecodeParameter("PLMNListforQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return PLMNListforQMC{}, 0, err
	}
	return v, len(data), err
}
