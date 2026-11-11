// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AdditionalULI represents ASN.1 AdditionalULI.
type AdditionalULI struct {
	NRCGI             NRCGI
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []AdditionalULIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalULI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AdditionalULI", v) }

// DecodeAdditionalULI decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalULI(data []byte) (AdditionalULI, int, error) {
	var v AdditionalULI
	err := ngapcodec.DecodeParameter("AdditionalULI", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalULI{}, 0, err
	}
	return v, len(data), err
}
