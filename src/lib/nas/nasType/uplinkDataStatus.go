package nasType

// IEUplinkDataStatus identifies the optional UplinkDataStatus IE on the wire.
const IEUplinkDataStatus uint8 = 0x40

// UplinkDataStatus is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type UplinkDataStatus struct {
	Value []byte
}
