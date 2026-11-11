package nasType

// IENASMessageContainer identifies the optional NASMessageContainer IE on the wire.
const IENASMessageContainer uint8 = 0x71

// NASMessageContainer is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type NASMessageContainer struct {
	Value []byte
}
