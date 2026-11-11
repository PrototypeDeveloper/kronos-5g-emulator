// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFRegionID represents ASN.1 AMFRegionID.
type AMFRegionID struct {
	Value BitString `aper:"sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFRegionID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AMFRegionID", v) }

// DecodeAMFRegionID decodes a standalone APER value and reports the consumed octets.
func DecodeAMFRegionID(data []byte) (AMFRegionID, int, error) {
	var v AMFRegionID
	err := ngapcodec.DecodeParameter("AMFRegionID", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFRegionID{}, 0, err
	}
	return v, len(data), err
}
