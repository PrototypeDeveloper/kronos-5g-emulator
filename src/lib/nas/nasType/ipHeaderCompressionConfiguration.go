package nasType

// IEIPHeaderCompressionConfiguration identifies the optional IPHeaderCompressionConfiguration IE on the wire.
const IEIPHeaderCompressionConfiguration uint8 = 0x66

// IPHeaderCompressionConfiguration is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type IPHeaderCompressionConfiguration struct {
	Value []byte
}
