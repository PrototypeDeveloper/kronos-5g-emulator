package nasType

// IEConfiguredNSSAI identifies the optional ConfiguredNSSAI IE on the wire.
const IEConfiguredNSSAI uint8 = 0x31

// ConfiguredNSSAI is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type ConfiguredNSSAI struct {
	Value []byte
}
