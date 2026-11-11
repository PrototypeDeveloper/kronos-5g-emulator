package nasType

// IEAlwaysOnPDUSessionIndication identifies the optional AlwaysOnPDUSessionIndication IE on the wire.
const IEAlwaysOnPDUSessionIndication uint8 = 0x80

// AlwaysOnPDUSessionIndication is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type AlwaysOnPDUSessionIndication struct {
	Value uint8
}
