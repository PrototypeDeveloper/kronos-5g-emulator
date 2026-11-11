// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PNINPNBasedMDT represents ASN.1 PNI-NPNBasedMDT.
type PNINPNBasedMDT struct {
	CAGListforMDT     CAGListforMDT
	HasIEExtensions   bool
	IEExtensions      []PNINPNBasedMDTExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PNINPNBasedMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PNI-NPNBasedMDT", v)
}

// DecodePNINPNBasedMDT decodes a standalone APER value and reports the consumed octets.
func DecodePNINPNBasedMDT(data []byte) (PNINPNBasedMDT, int, error) {
	var v PNINPNBasedMDT
	err := ngapcodec.DecodeParameter("PNI-NPNBasedMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return PNINPNBasedMDT{}, 0, err
	}
	return v, len(data), err
}
