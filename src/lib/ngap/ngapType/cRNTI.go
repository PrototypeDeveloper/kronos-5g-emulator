// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CRNTI represents ASN.1 C-RNTI.
type CRNTI struct {
	Value BitString `aper:"sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CRNTI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("C-RNTI", v) }

// DecodeCRNTI decodes a standalone APER value and reports the consumed octets.
func DecodeCRNTI(data []byte) (CRNTI, int, error) {
	var v CRNTI
	err := ngapcodec.DecodeParameter("C-RNTI", data, &v)
	if ngapcodec.IsFatal(err) {
		return CRNTI{}, 0, err
	}
	return v, len(data), err
}
