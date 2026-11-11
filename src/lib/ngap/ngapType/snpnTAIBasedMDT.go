// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNTAIBasedMDT represents ASN.1 SNPN-TAIBasedMDT.
type SNPNTAIBasedMDT struct {
	SNPNTAIListforMDT SNPNTAIListforMDT
	HasIEExtensions   bool
	IEExtensions      []SNPNTAIBasedMDTExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNTAIBasedMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPN-TAIBasedMDT", v)
}

// DecodeSNPNTAIBasedMDT decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNTAIBasedMDT(data []byte) (SNPNTAIBasedMDT, int, error) {
	var v SNPNTAIBasedMDT
	err := ngapcodec.DecodeParameter("SNPN-TAIBasedMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNTAIBasedMDT{}, 0, err
	}
	return v, len(data), err
}
