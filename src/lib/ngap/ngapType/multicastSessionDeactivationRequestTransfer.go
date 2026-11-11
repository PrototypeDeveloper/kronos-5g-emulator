// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MulticastSessionDeactivationRequestTransfer represents ASN.1 MulticastSessionDeactivationRequestTransfer.
type MulticastSessionDeactivationRequestTransfer struct {
	MBSSessionID      MBSSessionID
	HasIEExtensions   bool
	IEExtensions      []MulticastSessionDeactivationRequestTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionDeactivationRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionDeactivationRequestTransfer", v)
}

// DecodeMulticastSessionDeactivationRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionDeactivationRequestTransfer(data []byte) (MulticastSessionDeactivationRequestTransfer, int, error) {
	var v MulticastSessionDeactivationRequestTransfer
	err := ngapcodec.DecodeParameter("MulticastSessionDeactivationRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionDeactivationRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
