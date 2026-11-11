package nasType

// IECause5GMM identifies the optional Cause5GMM IE on the wire.
const IECause5GMM uint8 = 0x58

// Cause5GMM is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type Cause5GMM struct {
	Value uint8
}
