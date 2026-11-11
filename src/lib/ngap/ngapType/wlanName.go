// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WLANName represents ASN.1 WLANName.
type WLANName struct {
	Value []byte `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WLANName) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("WLANName", v) }

// DecodeWLANName decodes a standalone APER value and reports the consumed octets.
func DecodeWLANName(data []byte) (WLANName, int, error) {
	var v WLANName
	err := ngapcodec.DecodeParameter("WLANName", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANName{}, 0, err
	}
	return v, len(data), err
}
