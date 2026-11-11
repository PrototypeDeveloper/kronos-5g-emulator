package nasType

// IET3448Value identifies the optional T3448Value IE on the wire.
const IET3448Value uint8 = 0x6b

// T3448Value is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type T3448Value struct {
	Value uint8
}
