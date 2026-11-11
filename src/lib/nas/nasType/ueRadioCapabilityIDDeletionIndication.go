package nasType

// IEUERadioCapabilityIDDeletionIndication identifies the optional UERadioCapabilityIDDeletionIndication IE on the wire.
const IEUERadioCapabilityIDDeletionIndication uint8 = 0xa0

// UERadioCapabilityIDDeletionIndication is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type UERadioCapabilityIDDeletionIndication struct {
	Value uint8
}
