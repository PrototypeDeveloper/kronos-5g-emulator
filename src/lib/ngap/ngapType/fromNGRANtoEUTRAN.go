// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FromNGRANtoEUTRAN represents ASN.1 FromNGRANtoEUTRAN.
type FromNGRANtoEUTRAN struct {
	SourceNGRANnodeID IntersystemSONNGRANnodeID
	TargeteNBID       IntersystemSONeNBID
	HasIEExtensions   bool
	IEExtensions      []FromNGRANtoEUTRANExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FromNGRANtoEUTRAN) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FromNGRANtoEUTRAN", v)
}

// DecodeFromNGRANtoEUTRAN decodes a standalone APER value and reports the consumed octets.
func DecodeFromNGRANtoEUTRAN(data []byte) (FromNGRANtoEUTRAN, int, error) {
	var v FromNGRANtoEUTRAN
	err := ngapcodec.DecodeParameter("FromNGRANtoEUTRAN", data, &v)
	if ngapcodec.IsFatal(err) {
		return FromNGRANtoEUTRAN{}, 0, err
	}
	return v, len(data), err
}
