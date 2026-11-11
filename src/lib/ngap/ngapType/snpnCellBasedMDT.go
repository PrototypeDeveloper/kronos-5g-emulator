// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNCellBasedMDT represents ASN.1 SNPN-CellBasedMDT.
type SNPNCellBasedMDT struct {
	SNPNcellIdListforMDT SNPNCellIdListforMDT
	HasIEExtensions      bool
	IEExtensions         []SNPNCellBasedMDTExtIEsEntry
	UnknownExtensions    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNCellBasedMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPN-CellBasedMDT", v)
}

// DecodeSNPNCellBasedMDT decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNCellBasedMDT(data []byte) (SNPNCellBasedMDT, int, error) {
	var v SNPNCellBasedMDT
	err := ngapcodec.DecodeParameter("SNPN-CellBasedMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNCellBasedMDT{}, 0, err
	}
	return v, len(data), err
}
