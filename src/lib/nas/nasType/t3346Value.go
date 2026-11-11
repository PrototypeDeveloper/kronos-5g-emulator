package nasType

// IET3346Value identifies the optional T3346Value IE on the wire.
const IET3346Value uint8 = 0x5f

// T3346Value is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type T3346Value struct {
	Value uint8
}
