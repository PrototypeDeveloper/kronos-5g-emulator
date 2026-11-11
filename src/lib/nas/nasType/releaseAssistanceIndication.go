package nasType

// IEReleaseAssistanceIndication identifies the optional ReleaseAssistanceIndication IE on the wire.
const IEReleaseAssistanceIndication uint8 = 0xf0

// ReleaseAssistanceIndication is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type ReleaseAssistanceIndication struct {
	Value uint8
}
