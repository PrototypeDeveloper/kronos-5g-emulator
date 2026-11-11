package nasType

// IERequestedNSSAI identifies the optional RequestedNSSAI IE on the wire.
const IERequestedNSSAI uint8 = 0x2f

// RequestedNSSAI is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type RequestedNSSAI struct {
	Value []byte
}
