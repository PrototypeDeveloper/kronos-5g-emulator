package nasType

// IET3324Value identifies the optional T3324Value IE on the wire.
const IET3324Value uint8 = 0x6a

// T3324Value is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type T3324Value struct {
	Value uint8
}
