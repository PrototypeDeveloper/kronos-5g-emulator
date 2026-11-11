package nasType

// IEConfigurationUpdateIndication identifies the optional ConfigurationUpdateIndication IE on the wire.
const IEConfigurationUpdateIndication uint8 = 0xd0

// ConfigurationUpdateIndication is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type ConfigurationUpdateIndication struct {
	Value uint8
}
