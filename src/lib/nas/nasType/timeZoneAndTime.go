package nasType

// IEUniversalTimeAndLocalTimeZone identifies the optional UniversalTimeAndLocalTimeZone IE on the wire.
const IEUniversalTimeAndLocalTimeZone uint8 = 0x47

// TimeZoneAndTime is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type TimeZoneAndTime struct {
	Value []byte
}
