package nasType

// IEFullNameForNetwork identifies the optional FullNameForNetwork IE on the wire.
const IEFullNameForNetwork uint8 = 0x43

// IEShortNameForNetwork identifies the optional ShortNameForNetwork IE on the wire.
const IEShortNameForNetwork uint8 = 0x45

// NetworkName is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type NetworkName struct {
	Value []byte
}
