package nasType

// IECIoTSmallDataContainer identifies the optional CIoTSmallDataContainer IE on the wire.
const IECIoTSmallDataContainer uint8 = 0x6f

// CIoTSmallDataContainer is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type CIoTSmallDataContainer struct {
	Value []byte
}
