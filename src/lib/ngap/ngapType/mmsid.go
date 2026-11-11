// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MMSID represents ASN.1 MMSID.
type MMSID struct {
	Value []byte `aper:"sizeLB:1,sizeUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MMSID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MMSID", v) }

// DecodeMMSID decodes a standalone APER value and reports the consumed octets.
func DecodeMMSID(data []byte) (MMSID, int, error) {
	var v MMSID
	err := ngapcodec.DecodeParameter("MMSID", data, &v)
	if ngapcodec.IsFatal(err) {
		return MMSID{}, 0, err
	}
	return v, len(data), err
}
