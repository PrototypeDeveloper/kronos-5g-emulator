// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PLMNAreaBasedQMC represents ASN.1 PLMNAreaBasedQMC.
type PLMNAreaBasedQMC struct {
	PlmnListforQMC    PLMNListforQMC
	HasIEExtensions   bool
	IEExtensions      []PLMNAreaBasedQMCExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PLMNAreaBasedQMC) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PLMNAreaBasedQMC", v)
}

// DecodePLMNAreaBasedQMC decodes a standalone APER value and reports the consumed octets.
func DecodePLMNAreaBasedQMC(data []byte) (PLMNAreaBasedQMC, int, error) {
	var v PLMNAreaBasedQMC
	err := ngapcodec.DecodeParameter("PLMNAreaBasedQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return PLMNAreaBasedQMC{}, 0, err
	}
	return v, len(data), err
}
