// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupOrModRequestTransfer represents ASN.1 MBSSessionSetupOrModRequestTransfer.
type MBSSessionSetupOrModRequestTransfer struct {
	ProtocolIEs       []MBSSessionSetupOrModRequestTransferIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupOrModRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupOrModRequestTransfer", v)
}

// DecodeMBSSessionSetupOrModRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupOrModRequestTransfer(data []byte) (MBSSessionSetupOrModRequestTransfer, int, error) {
	var v MBSSessionSetupOrModRequestTransfer
	err := ngapcodec.DecodeParameter("MBSSessionSetupOrModRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupOrModRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
