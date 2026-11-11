package nasType

// IESORTransparentContainer identifies the optional SORTransparentContainer IE on the wire.
const IESORTransparentContainer uint8 = 0x73

// SORTransparentContainer is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type SORTransparentContainer struct {
	Value []byte
}
