package nasType

// IEIMEISV identifies the optional IMEISV IE on the wire.
const IEIMEISV uint8 = 0x77

// IMEISV is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type IMEISV struct {
	Value []byte
}
