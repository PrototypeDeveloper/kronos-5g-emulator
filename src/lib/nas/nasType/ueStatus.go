package nasType

// IEUEStatus identifies the optional UEStatus IE on the wire.
const IEUEStatus uint8 = 0x2b

// UEStatus is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type UEStatus struct {
	Value uint8
}
