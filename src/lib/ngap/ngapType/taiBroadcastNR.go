// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIBroadcastNR represents ASN.1 TAIBroadcastNR.
type TAIBroadcastNR struct {
	List []TAIBroadcastNRItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBroadcastNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIBroadcastNR", v)
}

// DecodeTAIBroadcastNR decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBroadcastNR(data []byte) (TAIBroadcastNR, int, error) {
	var v TAIBroadcastNR
	err := ngapcodec.DecodeParameter("TAIBroadcastNR", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBroadcastNR{}, 0, err
	}
	return v, len(data), err
}
