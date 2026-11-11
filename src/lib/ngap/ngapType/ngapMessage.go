// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGAPMessage represents ASN.1 NGAP-Message.
type NGAPMessage struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGAPMessage) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NGAP-Message", v) }

// DecodeNGAPMessage decodes a standalone APER value and reports the consumed octets.
func DecodeNGAPMessage(data []byte) (NGAPMessage, int, error) {
	var v NGAPMessage
	err := ngapcodec.DecodeParameter("NGAP-Message", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPMessage{}, 0, err
	}
	return v, len(data), err
}
