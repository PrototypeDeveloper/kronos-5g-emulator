// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ResetType represents ASN.1 ResetType.
type ResetTypePresent uint32

const (
	ResetTypePresentNothing ResetTypePresent = iota
	ResetTypePresentNGInterface
	ResetTypePresentPartOfNGInterface
	ResetTypePresentChoiceExtensions
)

type ResetType struct {
	Present           ResetTypePresent
	NGInterface       ResetAll
	PartOfNGInterface UEAssociatedLogicalNGConnectionList
	ChoiceExtensions  ResetTypeExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ResetType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ResetType", v) }

// DecodeResetType decodes a standalone APER value and reports the consumed octets.
func DecodeResetType(data []byte) (ResetType, int, error) {
	var v ResetType
	err := ngapcodec.DecodeParameter("ResetType", data, &v)
	if ngapcodec.IsFatal(err) {
		return ResetType{}, 0, err
	}
	return v, len(data), err
}
