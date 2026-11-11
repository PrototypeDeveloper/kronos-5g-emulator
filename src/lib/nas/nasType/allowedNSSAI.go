package nasType

// IEAllowedNSSAI identifies the optional AllowedNSSAI IE on the wire.
const IEAllowedNSSAI uint8 = 0x15

// AllowedNSSAI is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type AllowedNSSAI struct {
	Value []byte
}
