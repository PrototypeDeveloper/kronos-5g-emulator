package nasType

// IEAllowedPDUSessionStatus identifies the optional AllowedPDUSessionStatus IE on the wire.
const IEAllowedPDUSessionStatus uint8 = 0x25

// AllowedPDUSessionStatus is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type AllowedPDUSessionStatus struct {
	Value []byte
}
