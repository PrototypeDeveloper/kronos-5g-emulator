package nasType

// IEQoSRules identifies the optional QoSRules IE on the wire.
const IEQoSRules uint8 = 0x7a

// QoSRules is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type QoSRules struct {
	Value []byte
}
