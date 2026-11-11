package nasType

// IECapability5GSM identifies the optional Capability5GSM IE on the wire.
const IECapability5GSM uint8 = 0x28

// Capability5GSM is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type Capability5GSM struct {
	Value []byte
}
