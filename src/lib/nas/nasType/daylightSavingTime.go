package nasType

// IENetworkDaylightSavingTime identifies the optional NetworkDaylightSavingTime IE on the wire.
const IENetworkDaylightSavingTime uint8 = 0x49

// DaylightSavingTime is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type DaylightSavingTime struct {
	Value uint8
}
