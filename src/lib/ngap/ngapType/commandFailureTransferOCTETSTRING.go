// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CommandFailureTransferOCTETSTRING represents ASN.1 CommandFailureTransfer-OCTET-STRING.
type CommandFailureTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandFailureTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandFailureTransfer-OCTET-STRING", v)
}

// DecodeCommandFailureTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeCommandFailureTransferOCTETSTRING(data []byte) (CommandFailureTransferOCTETSTRING, int, error) {
	var v CommandFailureTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("CommandFailureTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandFailureTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
