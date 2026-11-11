// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastSessionUpdateRequestTransfer represents ASN.1 MulticastSessionUpdateRequestTransfer.
type MulticastSessionUpdateRequestTransfer struct {
	ProtocolIEs       []MulticastSessionUpdateRequestTransferIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionUpdateRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionUpdateRequestTransfer", v)
}

// DecodeMulticastSessionUpdateRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionUpdateRequestTransfer(data []byte) (MulticastSessionUpdateRequestTransfer, int, error) {
	var v MulticastSessionUpdateRequestTransfer
	err := ngapcodec.DecodeParameter("MulticastSessionUpdateRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionUpdateRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
