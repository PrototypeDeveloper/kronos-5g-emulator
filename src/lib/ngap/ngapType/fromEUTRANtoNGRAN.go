// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FromEUTRANtoNGRAN represents ASN.1 FromEUTRANtoNGRAN.
type FromEUTRANtoNGRAN struct {
	SourceeNBID       IntersystemSONeNBID
	TargetNGRANnodeID IntersystemSONNGRANnodeID
	HasIEExtensions   bool
	IEExtensions      []FromEUTRANtoNGRANExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FromEUTRANtoNGRAN) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FromEUTRANtoNGRAN", v)
}

// DecodeFromEUTRANtoNGRAN decodes a standalone APER value and reports the consumed octets.
func DecodeFromEUTRANtoNGRAN(data []byte) (FromEUTRANtoNGRAN, int, error) {
	var v FromEUTRANtoNGRAN
	err := ngapcodec.DecodeParameter("FromEUTRANtoNGRAN", data, &v)
	if ngapcodec.IsFatal(err) {
		return FromEUTRANtoNGRAN{}, 0, err
	}
	return v, len(data), err
}
