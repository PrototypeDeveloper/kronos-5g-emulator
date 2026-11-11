// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CommandRequestTransfer represents ASN.1 CommandRequestTransfer.
type CommandRequestTransfer struct {
	ProtocolIEs       []CommandRequestTransferIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandRequestTransfer", v)
}

// DecodeCommandRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeCommandRequestTransfer(data []byte) (CommandRequestTransfer, int, error) {
	var v CommandRequestTransfer
	err := ngapcodec.DecodeParameter("CommandRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
