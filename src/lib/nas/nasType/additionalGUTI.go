package nasType

// IEAdditionalGUTI identifies the optional AdditionalGUTI IE on the wire.
const IEAdditionalGUTI uint8 = 0x77

// AdditionalGUTI is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type AdditionalGUTI struct {
	Value []byte
}
