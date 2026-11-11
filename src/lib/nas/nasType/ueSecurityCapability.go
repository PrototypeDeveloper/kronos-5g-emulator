package nasType

// IEUESecurityCapability identifies the optional UESecurityCapability IE on the wire.
const IEUESecurityCapability uint8 = 0x2e

// UESecurityCapability is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type UESecurityCapability struct {
	Value []byte
}
