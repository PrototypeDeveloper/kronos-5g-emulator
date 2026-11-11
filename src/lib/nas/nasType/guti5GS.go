package nasType

// IEGUTI5GS identifies the optional GUTI5GS IE on the wire.
const IEGUTI5GS uint8 = 0x77

// GUTI5GS is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type GUTI5GS struct {
	Value []byte
}
