package nasType

// IEUERadioCapabilityID identifies the optional UERadioCapabilityID IE on the wire.
const IEUERadioCapabilityID uint8 = 0x67

// UERadioCapabilityID is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type UERadioCapabilityID struct {
	Value []byte
}
