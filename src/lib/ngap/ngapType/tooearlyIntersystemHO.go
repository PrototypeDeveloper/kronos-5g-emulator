// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TooearlyIntersystemHO represents ASN.1 TooearlyIntersystemHO.
type TooearlyIntersystemHO struct {
	SourcecellID            EUTRACGI
	FailurecellID           NGRANCGI
	HasUERLFReportContainer bool
	UERLFReportContainer    UERLFReportContainer
	HasIEExtensions         bool
	IEExtensions            []TooearlyIntersystemHOExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TooearlyIntersystemHO) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TooearlyIntersystemHO", v)
}

// DecodeTooearlyIntersystemHO decodes a standalone APER value and reports the consumed octets.
func DecodeTooearlyIntersystemHO(data []byte) (TooearlyIntersystemHO, int, error) {
	var v TooearlyIntersystemHO
	err := ngapcodec.DecodeParameter("TooearlyIntersystemHO", data, &v)
	if ngapcodec.IsFatal(err) {
		return TooearlyIntersystemHO{}, 0, err
	}
	return v, len(data), err
}
