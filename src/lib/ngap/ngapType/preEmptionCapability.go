// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PreEmptionCapability represents ASN.1 Pre-emptionCapability.
const (
	PreEmptionCapabilityPresentShallNotTriggerPreEmption Enumerated = 0
	PreEmptionCapabilityPresentMayTriggerPreEmption      Enumerated = 1
)

type PreEmptionCapability struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PreEmptionCapability) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Pre-emptionCapability", v)
}

// DecodePreEmptionCapability decodes a standalone APER value and reports the consumed octets.
func DecodePreEmptionCapability(data []byte) (PreEmptionCapability, int, error) {
	var v PreEmptionCapability
	err := ngapcodec.DecodeParameter("Pre-emptionCapability", data, &v)
	if ngapcodec.IsFatal(err) {
		return PreEmptionCapability{}, 0, err
	}
	return v, len(data), err
}
