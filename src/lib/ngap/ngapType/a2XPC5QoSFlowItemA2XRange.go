// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// A2XPC5QoSFlowItemA2XRange represents ASN.1 A2XPC5QoSFlowItemA2XRange.
const (
	A2XPC5QoSFlowItemA2XRangePresentM50   Enumerated = 0
	A2XPC5QoSFlowItemA2XRangePresentM80   Enumerated = 1
	A2XPC5QoSFlowItemA2XRangePresentM180  Enumerated = 2
	A2XPC5QoSFlowItemA2XRangePresentM200  Enumerated = 3
	A2XPC5QoSFlowItemA2XRangePresentM350  Enumerated = 4
	A2XPC5QoSFlowItemA2XRangePresentM400  Enumerated = 5
	A2XPC5QoSFlowItemA2XRangePresentM500  Enumerated = 6
	A2XPC5QoSFlowItemA2XRangePresentM700  Enumerated = 7
	A2XPC5QoSFlowItemA2XRangePresentM1000 Enumerated = 8
)

type A2XPC5QoSFlowItemA2XRange struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:8"`
}

// Encode encodes this parameter as a standalone APER value.
func (v A2XPC5QoSFlowItemA2XRange) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:A2XPC5QoSFlowItemA2XRange", v)
}

// DecodeA2XPC5QoSFlowItemA2XRange decodes a standalone APER value and reports consumed octets.
func DecodeA2XPC5QoSFlowItemA2XRange(data []byte) (A2XPC5QoSFlowItemA2XRange, int, error) {
	var v A2XPC5QoSFlowItemA2XRange
	err := ngapcodec.DecodeParameter("inline:A2XPC5QoSFlowItemA2XRange", data, &v)
	if ngapcodec.IsFatal(err) {
		return A2XPC5QoSFlowItemA2XRange{}, 0, err
	}
	return v, len(data), err
}
