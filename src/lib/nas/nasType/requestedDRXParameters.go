package nasType

// IERequestedDRXParameters identifies the optional RequestedDRXParameters IE on the wire.
const IERequestedDRXParameters uint8 = 0x51

// RequestedDRXParameters is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type RequestedDRXParameters struct {
	Value uint8
}
