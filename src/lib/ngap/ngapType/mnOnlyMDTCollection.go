// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MNOnlyMDTCollection represents ASN.1 MN-only-MDT-collection.
const (
	MNOnlyMDTCollectionPresentMnOnly Enumerated = 0
)

type MNOnlyMDTCollection struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MNOnlyMDTCollection) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MN-only-MDT-collection", v)
}

// DecodeMNOnlyMDTCollection decodes a standalone APER value and reports the consumed octets.
func DecodeMNOnlyMDTCollection(data []byte) (MNOnlyMDTCollection, int, error) {
	var v MNOnlyMDTCollection
	err := ngapcodec.DecodeParameter("MN-only-MDT-collection", data, &v)
	if ngapcodec.IsFatal(err) {
		return MNOnlyMDTCollection{}, 0, err
	}
	return v, len(data), err
}
