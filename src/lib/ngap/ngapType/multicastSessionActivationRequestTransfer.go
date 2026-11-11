// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastSessionActivationRequestTransfer represents ASN.1 MulticastSessionActivationRequestTransfer.
type MulticastSessionActivationRequestTransfer struct {
	MBSSessionID      MBSSessionID
	HasIEExtensions   bool
	IEExtensions      []MulticastSessionActivationRequestTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionActivationRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionActivationRequestTransfer", v)
}

// DecodeMulticastSessionActivationRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionActivationRequestTransfer(data []byte) (MulticastSessionActivationRequestTransfer, int, error) {
	var v MulticastSessionActivationRequestTransfer
	err := ngapcodec.DecodeParameter("MulticastSessionActivationRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionActivationRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
