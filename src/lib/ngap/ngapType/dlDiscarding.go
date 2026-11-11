// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DLDiscarding represents ASN.1 DLDiscarding.
type DLDiscarding struct {
	DRBsSubjectToDLDiscarding DRBsSubjectToDLDiscardingList
	HasIEExtension            bool
	IEExtension               []DLDiscardingExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLDiscarding) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DLDiscarding", v) }

// DecodeDLDiscarding decodes a standalone APER value and reports the consumed octets.
func DecodeDLDiscarding(data []byte) (DLDiscarding, int, error) {
	var v DLDiscarding
	err := ngapcodec.DecodeParameter("DLDiscarding", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLDiscarding{}, 0, err
	}
	return v, len(data), err
}
