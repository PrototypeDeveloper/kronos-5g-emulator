package nasType

// IESNSSAI identifies the optional SNSSAI IE on the wire.
const IESNSSAI uint8 = 0x22

// SNSSAI is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type SNSSAI struct {
	Value []byte
}
