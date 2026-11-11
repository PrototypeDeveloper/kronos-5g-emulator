package nasType

// IECapability5GMM identifies the optional Capability5GMM IE on the wire.
const IECapability5GMM uint8 = 0x10

// Capability5GMM is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type Capability5GMM struct {
	Value []byte
}
