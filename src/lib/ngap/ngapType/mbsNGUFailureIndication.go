// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSNGUFailureIndication represents ASN.1 MBS-NGUFailureIndication.
type MBSNGUFailureIndicationPresent uint32

const (
	MBSNGUFailureIndicationPresentNothing MBSNGUFailureIndicationPresent = iota
	MBSNGUFailureIndicationPresentLocationindependent
	MBSNGUFailureIndicationPresentLocationdependent
	MBSNGUFailureIndicationPresentChoiceExtensions
)

type MBSNGUFailureIndication struct {
	Present             MBSNGUFailureIndicationPresent
	Locationindependent MBSUPFailureIndication
	Locationdependent   MBSUPFailureIndicationList
	ChoiceExtensions    MBSNGUFailureIndicationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSNGUFailureIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-NGUFailureIndication", v)
}

// DecodeMBSNGUFailureIndication decodes a standalone APER value and reports the consumed octets.
func DecodeMBSNGUFailureIndication(data []byte) (MBSNGUFailureIndication, int, error) {
	var v MBSNGUFailureIndication
	err := ngapcodec.DecodeParameter("MBS-NGUFailureIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSNGUFailureIndication{}, 0, err
	}
	return v, len(data), err
}
