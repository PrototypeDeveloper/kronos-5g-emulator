// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFSetID represents ASN.1 AMFSetID.
type AMFSetID struct {
	Value BitString `aper:"sizeLB:10,sizeUB:10"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFSetID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AMFSetID", v) }

// DecodeAMFSetID decodes a standalone APER value and reports the consumed octets.
func DecodeAMFSetID(data []byte) (AMFSetID, int, error) {
	var v AMFSetID
	err := ngapcodec.DecodeParameter("AMFSetID", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFSetID{}, 0, err
	}
	return v, len(data), err
}
