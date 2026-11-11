package nasType

// IEEAPMessage identifies the optional EAPMessage IE on the wire.
const IEEAPMessage uint8 = 0x78

// EAPMessage is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type EAPMessage struct {
	Value []byte
}
