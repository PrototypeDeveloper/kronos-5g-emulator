package nasType

// IEBackOffTimerValue identifies the optional BackOffTimerValue IE on the wire.
const IEBackOffTimerValue uint8 = 0x37

// BackOffTimerValue is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type BackOffTimerValue struct {
	Value uint8
}
