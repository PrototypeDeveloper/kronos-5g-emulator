// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CNsubgroupID represents ASN.1 CNsubgroupID.
type CNsubgroupID struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CNsubgroupID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CNsubgroupID", v) }

// DecodeCNsubgroupID decodes a standalone APER value and reports the consumed octets.
func DecodeCNsubgroupID(data []byte) (CNsubgroupID, int, error) {
	var v CNsubgroupID
	err := ngapcodec.DecodeParameter("CNsubgroupID", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNsubgroupID{}, 0, err
	}
	return v, len(data), err
}
