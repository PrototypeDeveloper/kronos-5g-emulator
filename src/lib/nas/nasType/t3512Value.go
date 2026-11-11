package nasType

// IET3512Value identifies the optional T3512Value IE on the wire.
const IET3512Value uint8 = 0x5e

// T3512Value is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type T3512Value struct {
	Value uint8
}
