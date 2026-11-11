package nasType

// IEAdditionalConfigurationIndication identifies the optional AdditionalConfigurationIndication IE on the wire.
const IEAdditionalConfigurationIndication uint8 = 0xc0

// AdditionalConfigurationIndication is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type AdditionalConfigurationIndication struct {
	Value uint8
}
