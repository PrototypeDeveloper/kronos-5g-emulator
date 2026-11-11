// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemSONTransferType represents ASN.1 IntersystemSONTransferType.
type IntersystemSONTransferTypePresent uint32

const (
	IntersystemSONTransferTypePresentNothing IntersystemSONTransferTypePresent = iota
	IntersystemSONTransferTypePresentFromEUTRANtoNGRAN
	IntersystemSONTransferTypePresentFromNGRANtoEUTRAN
	IntersystemSONTransferTypePresentChoiceExtensions
)

type IntersystemSONTransferType struct {
	Present           IntersystemSONTransferTypePresent
	FromEUTRANtoNGRAN FromEUTRANtoNGRAN
	FromNGRANtoEUTRAN FromNGRANtoEUTRAN
	ChoiceExtensions  IntersystemSONTransferTypeExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONTransferType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONTransferType", v)
}

// DecodeIntersystemSONTransferType decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONTransferType(data []byte) (IntersystemSONTransferType, int, error) {
	var v IntersystemSONTransferType
	err := ngapcodec.DecodeParameter("IntersystemSONTransferType", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONTransferType{}, 0, err
	}
	return v, len(data), err
}
