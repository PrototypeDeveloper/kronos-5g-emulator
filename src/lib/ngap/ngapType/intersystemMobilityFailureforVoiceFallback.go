// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemMobilityFailureforVoiceFallback represents ASN.1 IntersystemMobilityFailureforVoiceFallback.
type IntersystemMobilityFailureforVoiceFallback struct {
	SourcecellID            NGRANCGI
	TargetcellID            EUTRACGI
	HasReconnectCellID      bool
	ReconnectCellID         EUTRACGI
	HasUERLFReportContainer bool
	UERLFReportContainer    UERLFReportContainer
	HasIEExtensions         bool
	IEExtensions            []IntersystemMobilityFailureforVoiceFallbackExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemMobilityFailureforVoiceFallback) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemMobilityFailureforVoiceFallback", v)
}

// DecodeIntersystemMobilityFailureforVoiceFallback decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemMobilityFailureforVoiceFallback(data []byte) (IntersystemMobilityFailureforVoiceFallback, int, error) {
	var v IntersystemMobilityFailureforVoiceFallback
	err := ngapcodec.DecodeParameter("IntersystemMobilityFailureforVoiceFallback", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemMobilityFailureforVoiceFallback{}, 0, err
	}
	return v, len(data), err
}
