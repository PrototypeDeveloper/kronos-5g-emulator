// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PeriodicityBound represents ASN.1 PeriodicityBound.
type PeriodicityBound struct {
	PeriodicityLowerBound Periodicity
	PeriodicityUpperBound Periodicity
	HasIEExtensions       bool
	IEExtensions          []PeriodicityBoundExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PeriodicityBound) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PeriodicityBound", v)
}

// DecodePeriodicityBound decodes a standalone APER value and reports the consumed octets.
func DecodePeriodicityBound(data []byte) (PeriodicityBound, int, error) {
	var v PeriodicityBound
	err := ngapcodec.DecodeParameter("PeriodicityBound", data, &v)
	if ngapcodec.IsFatal(err) {
		return PeriodicityBound{}, 0, err
	}
	return v, len(data), err
}
