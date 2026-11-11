package nasType

// IEPayloadContainer identifies the optional PayloadContainer IE on the wire.
const IEPayloadContainer uint8 = 0x7b

// PayloadContainer is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type PayloadContainer struct {
	Value []byte
}
