package nasType

// IEUEUsageSetting identifies the optional UEUsageSetting IE on the wire.
const IEUEUsageSetting uint8 = 0x18

// UEUsageSetting is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type UEUsageSetting struct {
	Value uint8
}
