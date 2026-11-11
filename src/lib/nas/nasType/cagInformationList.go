package nasType

// IECAGInformationList identifies the optional CAGInformationList IE on the wire.
const IECAGInformationList uint8 = 0x75

// CAGInformationList is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type CAGInformationList struct {
	Value []byte
}
