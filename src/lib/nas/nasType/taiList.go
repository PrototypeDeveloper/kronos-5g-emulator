package nasType

// IETAIList identifies the optional TAIList IE on the wire.
const IETAIList uint8 = 0x54

// TAIList is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type TAIList struct {
	Value []byte
}
