// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CommandResponseTransferOCTETSTRING represents ASN.1 CommandResponseTransfer-OCTET-STRING.
type CommandResponseTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandResponseTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandResponseTransfer-OCTET-STRING", v)
}

// DecodeCommandResponseTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeCommandResponseTransferOCTETSTRING(data []byte) (CommandResponseTransferOCTETSTRING, int, error) {
	var v CommandResponseTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("CommandResponseTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandResponseTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
