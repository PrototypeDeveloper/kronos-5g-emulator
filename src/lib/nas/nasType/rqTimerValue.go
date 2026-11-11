package nasType

// IERQTimerValue identifies the optional RQTimerValue IE on the wire.
const IERQTimerValue uint8 = 0x56

// RQTimerValue is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type RQTimerValue struct {
	Value uint8
}
