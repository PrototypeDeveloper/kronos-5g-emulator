// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CriticalityDiagnostics represents ASN.1 CriticalityDiagnostics.
type CriticalityDiagnostics struct {
	HasProcedureCode             bool
	ProcedureCode                ProcedureCode
	HasTriggeringMessage         bool
	TriggeringMessage            TriggeringMessage
	HasProcedureCriticality      bool
	ProcedureCriticality         Criticality
	HasIEsCriticalityDiagnostics bool
	IEsCriticalityDiagnostics    CriticalityDiagnosticsIEList
	HasIEExtensions              bool
	IEExtensions                 []CriticalityDiagnosticsExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CriticalityDiagnostics) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CriticalityDiagnostics", v)
}

// DecodeCriticalityDiagnostics decodes a standalone APER value and reports the consumed octets.
func DecodeCriticalityDiagnostics(data []byte) (CriticalityDiagnostics, int, error) {
	var v CriticalityDiagnostics
	err := ngapcodec.DecodeParameter("CriticalityDiagnostics", data, &v)
	if ngapcodec.IsFatal(err) {
		return CriticalityDiagnostics{}, 0, err
	}
	return v, len(data), err
}
