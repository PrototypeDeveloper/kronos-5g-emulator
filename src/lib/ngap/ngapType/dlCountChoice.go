// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DLCountChoice represents ASN.1 DLCountChoice.
type DLCountChoicePresent uint32

const (
	DLCountChoicePresentNothing DLCountChoicePresent = iota
	DLCountChoicePresentCount12bits
	DLCountChoicePresentCount18bits
	DLCountChoicePresentChoiceExtension
)

type DLCountChoice struct {
	Present         DLCountChoicePresent
	Count12bits     COUNTValueForPDCPSN12
	Count18bits     COUNTValueForPDCPSN18
	ChoiceExtension DLCountChoiceExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLCountChoice) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DLCountChoice", v) }

// DecodeDLCountChoice decodes a standalone APER value and reports the consumed octets.
func DecodeDLCountChoice(data []byte) (DLCountChoice, int, error) {
	var v DLCountChoice
	err := ngapcodec.DecodeParameter("DLCountChoice", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLCountChoice{}, 0, err
	}
	return v, len(data), err
}
