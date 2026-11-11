// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastTransportRequestTransfer represents ASN.1 BroadcastTransportRequestTransfer.
type BroadcastTransportRequestTransfer struct {
	MBSSessionID           MBSSessionID
	MBSSessionTNLInfoNGRAN MBSSessionTNLInfoNGRAN
	HasIEExtensions        bool
	IEExtensions           []BroadcastTransportRequestTransferExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportRequestTransfer", v)
}

// DecodeBroadcastTransportRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportRequestTransfer(data []byte) (BroadcastTransportRequestTransfer, int, error) {
	var v BroadcastTransportRequestTransfer
	err := ngapcodec.DecodeParameter("BroadcastTransportRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
