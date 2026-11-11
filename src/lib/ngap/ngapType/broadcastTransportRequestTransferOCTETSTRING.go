// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastTransportRequestTransferOCTETSTRING represents ASN.1 BroadcastTransportRequestTransfer-OCTET-STRING.
type BroadcastTransportRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportRequestTransfer-OCTET-STRING", v)
}

// DecodeBroadcastTransportRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportRequestTransferOCTETSTRING(data []byte) (BroadcastTransportRequestTransferOCTETSTRING, int, error) {
	var v BroadcastTransportRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("BroadcastTransportRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
