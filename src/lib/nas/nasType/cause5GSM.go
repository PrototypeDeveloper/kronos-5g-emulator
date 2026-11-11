package nasType

// IECause5GSM identifies the optional Cause5GSM IE on the wire.
const IECause5GSM uint8 = 0x59

// Cause5GSM is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type Cause5GSM struct {
	Value uint8
}
