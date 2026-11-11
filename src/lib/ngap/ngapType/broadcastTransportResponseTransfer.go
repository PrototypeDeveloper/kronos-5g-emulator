// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastTransportResponseTransfer represents ASN.1 BroadcastTransportResponseTransfer.
type BroadcastTransportResponseTransfer struct {
	MBSSessionID      MBSSessionID
	HasIEExtensions   bool
	IEExtensions      []BroadcastTransportResponseTransferExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportResponseTransfer", v)
}

// DecodeBroadcastTransportResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportResponseTransfer(data []byte) (BroadcastTransportResponseTransfer, int, error) {
	var v BroadcastTransportResponseTransfer
	err := ngapcodec.DecodeParameter("BroadcastTransportResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
