package nasType

// IEDNN identifies the optional DNN IE on the wire.
const IEDNN uint8 = 0x25

// DNN is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type DNN struct {
	Value []byte
}
