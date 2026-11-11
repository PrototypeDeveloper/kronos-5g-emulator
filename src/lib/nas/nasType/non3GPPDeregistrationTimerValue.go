package nasType

// IENon3GPPDeregistrationTimerValue identifies the optional Non3GPPDeregistrationTimerValue IE on the wire.
const IENon3GPPDeregistrationTimerValue uint8 = 0x5d

// Non3GPPDeregistrationTimerValue is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type Non3GPPDeregistrationTimerValue struct {
	Value uint8
}
