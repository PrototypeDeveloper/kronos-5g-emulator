// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SCTPTLAs represents ASN.1 SCTP-TLAs.
type SCTPTLAs struct {
	List []TransportLayerAddress `aper:"sizeLB:1,sizeUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SCTPTLAs) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("SCTP-TLAs", v) }

// DecodeSCTPTLAs decodes a standalone APER value and reports the consumed octets.
func DecodeSCTPTLAs(data []byte) (SCTPTLAs, int, error) {
	var v SCTPTLAs
	err := ngapcodec.DecodeParameter("SCTP-TLAs", data, &v)
	if ngapcodec.IsFatal(err) {
		return SCTPTLAs{}, 0, err
	}
	return v, len(data), err
}
