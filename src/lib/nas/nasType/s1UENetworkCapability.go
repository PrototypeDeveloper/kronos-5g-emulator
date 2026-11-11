package nasType

// IES1UENetworkCapability identifies the optional S1UENetworkCapability IE on the wire.
const IES1UENetworkCapability uint8 = 0x17

// S1UENetworkCapability is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type S1UENetworkCapability struct {
	Value []byte
}
