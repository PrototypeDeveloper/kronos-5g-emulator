// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CriticalityDiagnosticsIEItem represents ASN.1 CriticalityDiagnostics-IE-Item.
type CriticalityDiagnosticsIEItem struct {
	IECriticality     Criticality
	IEID              ProtocolIEID
	TypeOfError       TypeOfError
	HasIEExtensions   bool
	IEExtensions      []CriticalityDiagnosticsIEItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CriticalityDiagnosticsIEItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CriticalityDiagnostics-IE-Item", v)
}

// DecodeCriticalityDiagnosticsIEItem decodes a standalone APER value and reports the consumed octets.
func DecodeCriticalityDiagnosticsIEItem(data []byte) (CriticalityDiagnosticsIEItem, int, error) {
	var v CriticalityDiagnosticsIEItem
	err := ngapcodec.DecodeParameter("CriticalityDiagnostics-IE-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return CriticalityDiagnosticsIEItem{}, 0, err
	}
	return v, len(data), err
}
