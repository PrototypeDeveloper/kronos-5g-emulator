// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ProcedureStageChoiceExtIEsEntryPresent uint32

const (
	ProcedureStageChoiceExtIEsEntryPresentNothing ProcedureStageChoiceExtIEsEntryPresent = iota
	ProcedureStageChoiceExtIEsEntryPresentDLDiscarding
	ProcedureStageChoiceExtIEsEntryPresentUnknown
)

type ProcedureStageChoiceExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ProcedureStageChoiceExtIEsEntryPresent
	DLDiscarding DLDiscarding `ngap:"id:421,criticality:ignore,presence:mandatory"`
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ProcedureStageChoiceExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ProcedureStageChoice-ExtIEsEntry", v)
}

// DecodeProcedureStageChoiceExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeProcedureStageChoiceExtIEsEntry(data []byte) (ProcedureStageChoiceExtIEsEntry, int, error) {
	var v ProcedureStageChoiceExtIEsEntry
	err := ngapcodec.DecodeParameter("ProcedureStageChoice-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ProcedureStageChoiceExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
