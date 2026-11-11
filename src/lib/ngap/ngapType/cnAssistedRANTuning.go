// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CNAssistedRANTuning represents ASN.1 CNAssistedRANTuning.
type CNAssistedRANTuning struct {
	HasExpectedUEBehaviour bool
	ExpectedUEBehaviour    ExpectedUEBehaviour
	HasIEExtensions        bool
	IEExtensions           []CNAssistedRANTuningExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CNAssistedRANTuning) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CNAssistedRANTuning", v)
}

// DecodeCNAssistedRANTuning decodes a standalone APER value and reports the consumed octets.
func DecodeCNAssistedRANTuning(data []byte) (CNAssistedRANTuning, int, error) {
	var v CNAssistedRANTuning
	err := ngapcodec.DecodeParameter("CNAssistedRANTuning", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNAssistedRANTuning{}, 0, err
	}
	return v, len(data), err
}
