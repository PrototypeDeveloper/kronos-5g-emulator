// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CriticalityDiagnosticsIEList represents ASN.1 CriticalityDiagnostics-IE-List.
type CriticalityDiagnosticsIEList struct {
	List []CriticalityDiagnosticsIEItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CriticalityDiagnosticsIEList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CriticalityDiagnostics-IE-List", v)
}

// DecodeCriticalityDiagnosticsIEList decodes a standalone APER value and reports the consumed octets.
func DecodeCriticalityDiagnosticsIEList(data []byte) (CriticalityDiagnosticsIEList, int, error) {
	var v CriticalityDiagnosticsIEList
	err := ngapcodec.DecodeParameter("CriticalityDiagnostics-IE-List", data, &v)
	if ngapcodec.IsFatal(err) {
		return CriticalityDiagnosticsIEList{}, 0, err
	}
	return v, len(data), err
}
