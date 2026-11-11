// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BroadcastTransportFailureTransferOCTETSTRING represents ASN.1 BroadcastTransportFailureTransfer-OCTET-STRING.
type BroadcastTransportFailureTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportFailureTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportFailureTransfer-OCTET-STRING", v)
}

// DecodeBroadcastTransportFailureTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportFailureTransferOCTETSTRING(data []byte) (BroadcastTransportFailureTransferOCTETSTRING, int, error) {
	var v BroadcastTransportFailureTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("BroadcastTransportFailureTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportFailureTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
