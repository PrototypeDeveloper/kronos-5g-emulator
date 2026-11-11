// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PNINPNAreaScopeofMDT represents ASN.1 PNI-NPN-AreaScopeofMDT.
type PNINPNAreaScopeofMDT struct {
	CAGListforMDT     CAGListforMDT
	HasIEExtensions   bool
	IEExtensions      []PNINPNAreaScopeofMDTExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PNINPNAreaScopeofMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PNI-NPN-AreaScopeofMDT", v)
}

// DecodePNINPNAreaScopeofMDT decodes a standalone APER value and reports the consumed octets.
func DecodePNINPNAreaScopeofMDT(data []byte) (PNINPNAreaScopeofMDT, int, error) {
	var v PNINPNAreaScopeofMDT
	err := ngapcodec.DecodeParameter("PNI-NPN-AreaScopeofMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return PNINPNAreaScopeofMDT{}, 0, err
	}
	return v, len(data), err
}
