package nasType

// IEMaximumNumberOfSupportedPacketFilters identifies the optional MaximumNumberOfSupportedPacketFilters IE on the wire.
const IEMaximumNumberOfSupportedPacketFilters uint8 = 0x55

// MaximumNumberOfSupportedPacketFilters is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type MaximumNumberOfSupportedPacketFilters struct {
	Value []byte
}
