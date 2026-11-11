package nasType

// IEAdditionalInformation identifies the optional AdditionalInformation IE on the wire.
const IEAdditionalInformation uint8 = 0x24

// AdditionalInformation is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type AdditionalInformation struct {
	Value []byte
}
