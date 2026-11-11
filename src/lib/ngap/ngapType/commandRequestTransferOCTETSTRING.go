// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CommandRequestTransferOCTETSTRING represents ASN.1 CommandRequestTransfer-OCTET-STRING.
type CommandRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandRequestTransfer-OCTET-STRING", v)
}

// DecodeCommandRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeCommandRequestTransferOCTETSTRING(data []byte) (CommandRequestTransferOCTETSTRING, int, error) {
	var v CommandRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("CommandRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
