package nasType

// IEMICOIndication identifies the optional MICOIndication IE on the wire.
const IEMICOIndication uint8 = 0xb0

// MICOIndication is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type MICOIndication struct {
	Value uint8
}
