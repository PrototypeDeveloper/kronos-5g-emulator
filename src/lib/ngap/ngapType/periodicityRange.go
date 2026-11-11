// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PeriodicityRange represents ASN.1 PeriodicityRange.
type PeriodicityRangePresent uint32

const (
	PeriodicityRangePresentNothing PeriodicityRangePresent = iota
	PeriodicityRangePresentPeriodicityBound
	PeriodicityRangePresentPeriodicityList
	PeriodicityRangePresentChoiceExtensions
)

type PeriodicityRange struct {
	Present          PeriodicityRangePresent
	PeriodicityBound PeriodicityBound
	PeriodicityList  AllowedPeriodicityList
	ChoiceExtensions PeriodicityRangeExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicityRange) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicityRange", v)
}

// DecodePeriodicityRange decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicityRange(data []byte) (PeriodicityRange, int, error) {
	var v PeriodicityRange
	err := ngapcodec.DecodeParameter("PeriodicityRange", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicityRange{}, 0, err
	}
	return v, len(data), err
}
