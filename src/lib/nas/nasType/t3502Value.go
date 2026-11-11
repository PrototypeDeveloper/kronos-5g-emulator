package nasType

// IET3502Value identifies the optional T3502Value IE on the wire.
const IET3502Value uint8 = 0x16

// T3502Value is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type T3502Value struct {
	Value uint8
}
