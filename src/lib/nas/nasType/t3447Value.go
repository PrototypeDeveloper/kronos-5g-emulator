package nasType

// IET3447Value identifies the optional T3447Value IE on the wire.
const IET3447Value uint8 = 0x6c

// T3447Value is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type T3447Value struct {
	Value uint8
}
