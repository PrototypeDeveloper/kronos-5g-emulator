// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANTSSScope represents ASN.1 RAN-TSSScope.
type RANTSSScopePresent uint32

const (
	RANTSSScopePresentNothing RANTSSScopePresent = iota
	RANTSSScopePresentRANNodeLevel
	RANTSSScopePresentCellListLevel
	RANTSSScopePresentChoiceExtensions
)

type RANTSSScope struct {
	Present          RANTSSScopePresent
	RANNodeLevel     GlobalGNBID
	CellListLevel    RANTSSCellList
	ChoiceExtensions RANTSSScopeExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANTSSScope) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("RAN-TSSScope", v) }

// DecodeRANTSSScope decodes a standalone APER value and reports the consumed octets.
func DecodeRANTSSScope(data []byte) (RANTSSScope, int, error) {
	var v RANTSSScope
	err := ngapcodec.DecodeParameter("RAN-TSSScope", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTSSScope{}, 0, err
	}
	return v, len(data), err
}
