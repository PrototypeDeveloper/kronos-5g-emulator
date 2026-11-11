package nasType

// IESessionAMBR identifies the optional SessionAMBR IE on the wire.
const IESessionAMBR uint8 = 0x2a

// SessionAMBR is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type SessionAMBR struct {
	Value []byte
}
