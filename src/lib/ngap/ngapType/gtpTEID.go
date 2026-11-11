// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GTPTEID represents ASN.1 GTP-TEID.
type GTPTEID struct {
	Value []byte `aper:"sizeLB:4,sizeUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GTPTEID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GTP-TEID", v) }

// DecodeGTPTEID decodes a standalone APER value and reports the consumed octets.
func DecodeGTPTEID(data []byte) (GTPTEID, int, error) {
	var v GTPTEID
	err := ngapcodec.DecodeParameter("GTP-TEID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GTPTEID{}, 0, err
	}
	return v, len(data), err
}
