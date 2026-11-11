package nasType

// IETruncated5GSTMSIConfiguration identifies the optional Truncated5GSTMSIConfiguration IE on the wire.
const IETruncated5GSTMSIConfiguration uint8 = 0x1b

// Truncated5GSTMSIConfiguration is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type Truncated5GSTMSIConfiguration struct {
	Value uint8
}
