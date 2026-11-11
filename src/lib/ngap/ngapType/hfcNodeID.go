// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HFCNodeID represents ASN.1 HFCNode-ID.
type HFCNodeID struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HFCNodeID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("HFCNode-ID", v) }

// DecodeHFCNodeID decodes a standalone APER value and reports the consumed octets.
func DecodeHFCNodeID(data []byte) (HFCNodeID, int, error) {
	var v HFCNodeID
	err := ngapcodec.DecodeParameter("HFCNode-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return HFCNodeID{}, 0, err
	}
	return v, len(data), err
}
