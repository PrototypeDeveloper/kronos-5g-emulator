package nasType

// IEServingPLMNRateControl identifies the optional ServingPLMNRateControl IE on the wire.
const IEServingPLMNRateControl uint8 = 0x1e

// ServingPLMNRateControl is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type ServingPLMNRateControl struct {
	Value []byte
}
