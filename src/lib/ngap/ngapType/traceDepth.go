// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TraceDepth represents ASN.1 TraceDepth.
const (
	TraceDepthPresentMinimum                               Enumerated = 0
	TraceDepthPresentMedium                                Enumerated = 1
	TraceDepthPresentMaximum                               Enumerated = 2
	TraceDepthPresentMinimumWithoutVendorSpecificExtension Enumerated = 3
	TraceDepthPresentMediumWithoutVendorSpecificExtension  Enumerated = 4
	TraceDepthPresentMaximumWithoutVendorSpecificExtension Enumerated = 5
	TraceDepthPresentMinimumOnlyVendorSpecificTraceRecord  Enumerated = 6
	TraceDepthPresentMediumOnlyVendorSpecificTraceRecord   Enumerated = 7
	TraceDepthPresentMaximumOnlyVendorSpecificTraceRecord  Enumerated = 8
)

type TraceDepth struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TraceDepth) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TraceDepth", v) }

// DecodeTraceDepth decodes a standalone APER value and reports the consumed octets.
func DecodeTraceDepth(data []byte) (TraceDepth, int, error) {
	var v TraceDepth
	err := ngapcodec.DecodeParameter("TraceDepth", data, &v)
	if ngapcodec.IsFatal(err) {
		return TraceDepth{}, 0, err
	}
	return v, len(data), err
}
