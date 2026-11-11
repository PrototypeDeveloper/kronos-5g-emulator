package nasType

// IEServiceAreaList identifies the optional ServiceAreaList IE on the wire.
const IEServiceAreaList uint8 = 0x27

// ServiceAreaList is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type ServiceAreaList struct {
	Value []byte
}
