// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTSessionReleaseCommandTransfer represents ASN.1 AIOTSessionReleaseCommandTransfer.
type AIOTSessionReleaseCommandTransfer struct {
	ProtocolIEs       []AIOTSessionReleaseCommandTransferIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseCommandTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseCommandTransfer", v)
}

// DecodeAIOTSessionReleaseCommandTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseCommandTransfer(data []byte) (AIOTSessionReleaseCommandTransfer, int, error) {
	var v AIOTSessionReleaseCommandTransfer
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseCommandTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseCommandTransfer{}, 0, err
	}
	return v, len(data), err
}
