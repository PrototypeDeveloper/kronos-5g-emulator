// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ProcedureStageChoice represents ASN.1 ProcedureStageChoice.
type ProcedureStageChoicePresent uint32

const (
	ProcedureStageChoicePresentNothing ProcedureStageChoicePresent = iota
	ProcedureStageChoicePresentFirstDlCount
	ProcedureStageChoicePresentChoiceExtensions
)

type ProcedureStageChoice struct {
	Present          ProcedureStageChoicePresent
	FirstDlCount     FirstDLCount
	ChoiceExtensions ProcedureStageChoiceExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ProcedureStageChoice) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ProcedureStageChoice", v)
}

// DecodeProcedureStageChoice decodes a standalone APER value and reports the consumed octets.
func DecodeProcedureStageChoice(data []byte) (ProcedureStageChoice, int, error) {
	var v ProcedureStageChoice
	err := ngapcodec.DecodeParameter("ProcedureStageChoice", data, &v)
	if ngapcodec.IsFatal(err) {
		return ProcedureStageChoice{}, 0, err
	}
	return v, len(data), err
}
