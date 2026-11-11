package nasType

// IECongestionReAttemptIndicator5GSM identifies the optional CongestionReAttemptIndicator5GSM IE on the wire.
const IECongestionReAttemptIndicator5GSM uint8 = 0x61

// CongestionReAttemptIndicator5GSM is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type CongestionReAttemptIndicator5GSM struct {
	Value uint8
}
