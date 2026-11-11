package nasType

// IEUpdateType5GS identifies the optional UpdateType5GS IE on the wire.
const IEUpdateType5GS uint8 = 0x53

// UpdateType5GS is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type UpdateType5GS struct {
	Value uint8
}
