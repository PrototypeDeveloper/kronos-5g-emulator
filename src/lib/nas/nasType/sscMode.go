package nasType

// IESSCMode identifies the optional SSCMode IE on the wire.
const IESSCMode uint8 = 0xa0

// SSCMode is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type SSCMode struct {
	Value uint8
}
