// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeUEStayedInCellEnhancedGranularity represents ASN.1 TimeUEStayedInCellEnhancedGranularity.
type TimeUEStayedInCellEnhancedGranularity struct {
	Value int64 `aper:"valueLB:0,valueUB:40950"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeUEStayedInCellEnhancedGranularity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TimeUEStayedInCellEnhancedGranularity", v)
}

// DecodeTimeUEStayedInCellEnhancedGranularity decodes a standalone APER value and reports the consumed octets.
func DecodeTimeUEStayedInCellEnhancedGranularity(data []byte) (TimeUEStayedInCellEnhancedGranularity, int, error) {
	var v TimeUEStayedInCellEnhancedGranularity
	err := ngapcodec.DecodeParameter("TimeUEStayedInCellEnhancedGranularity", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeUEStayedInCellEnhancedGranularity{}, 0, err
	}
	return v, len(data), err
}
