// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RRCContainer represents ASN.1 RRCContainer.
type RRCContainer struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RRCContainer) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("RRCContainer", v) }

// DecodeRRCContainer decodes a standalone APER value and reports the consumed octets.
func DecodeRRCContainer(data []byte) (RRCContainer, int, error) {
	var v RRCContainer
	err := ngapcodec.DecodeParameter("RRCContainer", data, &v)
	if ngapcodec.IsFatal(err) {
		return RRCContainer{}, 0, err
	}
	return v, len(data), err
}
