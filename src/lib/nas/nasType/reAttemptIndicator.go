package nasType

// IEReAttemptIndicator identifies the optional ReAttemptIndicator IE on the wire.
const IEReAttemptIndicator uint8 = 0x1d

// ReAttemptIndicator is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type ReAttemptIndicator struct {
	Value uint8
}
