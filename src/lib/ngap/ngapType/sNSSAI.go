// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNSSAI represents ASN.1 S-NSSAI.
type SNSSAI struct {
	SST               SST
	HasSD             bool
	SD                SD
	HasIEExtensions   bool
	IEExtensions      []SNSSAIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNSSAI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("S-NSSAI", v) }

// DecodeSNSSAI decodes a standalone APER value and reports the consumed octets.
func DecodeSNSSAI(data []byte) (SNSSAI, int, error) {
	var v SNSSAI
	err := ngapcodec.DecodeParameter("S-NSSAI", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNSSAI{}, 0, err
	}
	return v, len(data), err
}
