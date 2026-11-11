package nasType

// IESMSIndication identifies the optional SMSIndication IE on the wire.
const IESMSIndication uint8 = 0xf0

// SMSIndication is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type SMSIndication struct {
	Value uint8
}
