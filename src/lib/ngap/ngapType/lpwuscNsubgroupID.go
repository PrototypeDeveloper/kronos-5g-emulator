// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LPWUSCNsubgroupID represents ASN.1 LPWUSCNsubgroupID.
type LPWUSCNsubgroupID struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:30"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LPWUSCNsubgroupID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LPWUSCNsubgroupID", v)
}

// DecodeLPWUSCNsubgroupID decodes a standalone APER value and reports the consumed octets.
func DecodeLPWUSCNsubgroupID(data []byte) (LPWUSCNsubgroupID, int, error) {
	var v LPWUSCNsubgroupID
	err := ngapcodec.DecodeParameter("LPWUSCNsubgroupID", data, &v)
	if ngapcodec.IsFatal(err) {
		return LPWUSCNsubgroupID{}, 0, err
	}
	return v, len(data), err
}
