package nasType

// IEABBA identifies the optional ABBA IE on the wire.
const IEABBA uint8 = 0x38

// ABBA is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type ABBA struct {
	Value []byte
}
