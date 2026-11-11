package nasType

// IEEquivalentPLMNs identifies the optional EquivalentPLMNs IE on the wire.
const IEEquivalentPLMNs uint8 = 0x4a

// EquivalentPLMNs is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type EquivalentPLMNs struct {
	Value []byte
}
