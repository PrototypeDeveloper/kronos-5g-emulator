// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaScopeOfMDTEUTRA represents ASN.1 AreaScopeOfMDT-EUTRA.
type AreaScopeOfMDTEUTRAPresent uint32

const (
	AreaScopeOfMDTEUTRAPresentNothing AreaScopeOfMDTEUTRAPresent = iota
	AreaScopeOfMDTEUTRAPresentCellBased
	AreaScopeOfMDTEUTRAPresentTABased
	AreaScopeOfMDTEUTRAPresentPLMNWide
	AreaScopeOfMDTEUTRAPresentTAIBased
	AreaScopeOfMDTEUTRAPresentChoiceExtensions
)

type AreaScopeOfMDTEUTRA struct {
	Present          AreaScopeOfMDTEUTRAPresent
	CellBased        CellBasedMDTEUTRA
	TABased          TABasedMDT
	PLMNWide         struct{}
	TAIBased         TAIBasedMDT
	ChoiceExtensions AreaScopeOfMDTEUTRAExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfMDTEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfMDT-EUTRA", v)
}

// DecodeAreaScopeOfMDTEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfMDTEUTRA(data []byte) (AreaScopeOfMDTEUTRA, int, error) {
	var v AreaScopeOfMDTEUTRA
	err := ngapcodec.DecodeParameter("AreaScopeOfMDT-EUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfMDTEUTRA{}, 0, err
	}
	return v, len(data), err
}
