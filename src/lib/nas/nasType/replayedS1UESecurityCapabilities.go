package nasType

// IEReplayedS1UESecurityCapabilities identifies the optional ReplayedS1UESecurityCapabilities IE on the wire.
const IEReplayedS1UESecurityCapabilities uint8 = 0x19

// ReplayedS1UESecurityCapabilities is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type ReplayedS1UESecurityCapabilities struct {
	Value []byte
}
