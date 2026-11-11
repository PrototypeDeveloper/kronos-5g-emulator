package nasType

// IELocalTimeZone identifies the optional LocalTimeZone IE on the wire.
const IELocalTimeZone uint8 = 0x46

// TimeZone is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type TimeZone struct {
	Value uint8
}
