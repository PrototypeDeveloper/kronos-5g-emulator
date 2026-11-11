// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNBasedMDT represents ASN.1 SNPN-BasedMDT.
type SNPNBasedMDT struct {
	SNPNListforMDT    SNPNListforMDT
	HasIEExtensions   bool
	IEExtensions      []SNPNBasedMDTExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNBasedMDT) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SNPN-BasedMDT", v) }

// DecodeSNPNBasedMDT decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNBasedMDT(data []byte) (SNPNBasedMDT, int, error) {
	var v SNPNBasedMDT
	err := ngapcodec.DecodeParameter("SNPN-BasedMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNBasedMDT{}, 0, err
	}
	return v, len(data), err
}
