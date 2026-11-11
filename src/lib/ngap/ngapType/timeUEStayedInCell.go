// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeUEStayedInCell represents ASN.1 TimeUEStayedInCell.
type TimeUEStayedInCell struct {
	Value int64 `aper:"valueLB:0,valueUB:4095"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeUEStayedInCell) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TimeUEStayedInCell", v)
}

// DecodeTimeUEStayedInCell decodes a standalone APER value and reports the consumed octets.
func DecodeTimeUEStayedInCell(data []byte) (TimeUEStayedInCell, int, error) {
	var v TimeUEStayedInCell
	err := ngapcodec.DecodeParameter("TimeUEStayedInCell", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeUEStayedInCell{}, 0, err
	}
	return v, len(data), err
}
