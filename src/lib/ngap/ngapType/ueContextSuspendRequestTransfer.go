// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEContextSuspendRequestTransfer represents ASN.1 UEContextSuspendRequestTransfer.
type UEContextSuspendRequestTransfer struct {
	HasSuspendIndicator bool
	SuspendIndicator    SuspendIndicator
	HasIEExtensions     bool
	IEExtensions        []UEContextSuspendRequestTransferExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEContextSuspendRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEContextSuspendRequestTransfer", v)
}

// DecodeUEContextSuspendRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeUEContextSuspendRequestTransfer(data []byte) (UEContextSuspendRequestTransfer, int, error) {
	var v UEContextSuspendRequestTransfer
	err := ngapcodec.DecodeParameter("UEContextSuspendRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEContextSuspendRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
