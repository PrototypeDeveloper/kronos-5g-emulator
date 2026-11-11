// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TAIBroadcastEUTRA represents ASN.1 TAIBroadcastEUTRA.
type TAIBroadcastEUTRA struct {
	List []TAIBroadcastEUTRAItem `aper:"sizeLB:1,sizeUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBroadcastEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIBroadcastEUTRA", v)
}

// DecodeTAIBroadcastEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBroadcastEUTRA(data []byte) (TAIBroadcastEUTRA, int, error) {
	var v TAIBroadcastEUTRA
	err := ngapcodec.DecodeParameter("TAIBroadcastEUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBroadcastEUTRA{}, 0, err
	}
	return v, len(data), err
}
