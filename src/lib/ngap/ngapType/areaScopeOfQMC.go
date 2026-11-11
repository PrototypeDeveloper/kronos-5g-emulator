// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaScopeOfQMC represents ASN.1 AreaScopeOfQMC.
type AreaScopeOfQMCPresent uint32

const (
	AreaScopeOfQMCPresentNothing AreaScopeOfQMCPresent = iota
	AreaScopeOfQMCPresentCellBased
	AreaScopeOfQMCPresentTABased
	AreaScopeOfQMCPresentTAIBased
	AreaScopeOfQMCPresentPLMNAreaBased
	AreaScopeOfQMCPresentChoiceExtensions
)

type AreaScopeOfQMC struct {
	Present          AreaScopeOfQMCPresent
	CellBased        CellBasedQMC
	TABased          TABasedQMC
	TAIBased         TAIBasedQMC
	PLMNAreaBased    PLMNAreaBasedQMC
	ChoiceExtensions AreaScopeOfQMCExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfQMC) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfQMC", v)
}

// DecodeAreaScopeOfQMC decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfQMC(data []byte) (AreaScopeOfQMC, int, error) {
	var v AreaScopeOfQMC
	err := ngapcodec.DecodeParameter("AreaScopeOfQMC", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfQMC{}, 0, err
	}
	return v, len(data), err
}
