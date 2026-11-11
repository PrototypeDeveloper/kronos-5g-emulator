// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaScopeOfMDTNR represents ASN.1 AreaScopeOfMDT-NR.
type AreaScopeOfMDTNRPresent uint32

const (
	AreaScopeOfMDTNRPresentNothing AreaScopeOfMDTNRPresent = iota
	AreaScopeOfMDTNRPresentCellBased
	AreaScopeOfMDTNRPresentTABased
	AreaScopeOfMDTNRPresentPLMNWide
	AreaScopeOfMDTNRPresentTAIBased
	AreaScopeOfMDTNRPresentChoiceExtensions
)

type AreaScopeOfMDTNR struct {
	Present          AreaScopeOfMDTNRPresent
	CellBased        CellBasedMDTNR
	TABased          TABasedMDT
	PLMNWide         struct{}
	TAIBased         TAIBasedMDT
	ChoiceExtensions AreaScopeOfMDTNRExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfMDTNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfMDT-NR", v)
}

// DecodeAreaScopeOfMDTNR decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfMDTNR(data []byte) (AreaScopeOfMDTNR, int, error) {
	var v AreaScopeOfMDTNR
	err := ngapcodec.DecodeParameter("AreaScopeOfMDT-NR", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfMDTNR{}, 0, err
	}
	return v, len(data), err
}
