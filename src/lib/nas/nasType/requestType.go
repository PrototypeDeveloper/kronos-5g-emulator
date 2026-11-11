package nasType

// IERequestType identifies the optional RequestType IE on the wire.
const IERequestType uint8 = 0x80

// RequestType is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type RequestType struct {
	Value uint8
}
