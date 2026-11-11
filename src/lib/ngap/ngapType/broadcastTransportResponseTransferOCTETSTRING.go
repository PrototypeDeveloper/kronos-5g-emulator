// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastTransportResponseTransferOCTETSTRING represents ASN.1 BroadcastTransportResponseTransfer-OCTET-STRING.
type BroadcastTransportResponseTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportResponseTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportResponseTransfer-OCTET-STRING", v)
}

// DecodeBroadcastTransportResponseTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportResponseTransferOCTETSTRING(data []byte) (BroadcastTransportResponseTransferOCTETSTRING, int, error) {
	var v BroadcastTransportResponseTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("BroadcastTransportResponseTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportResponseTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
