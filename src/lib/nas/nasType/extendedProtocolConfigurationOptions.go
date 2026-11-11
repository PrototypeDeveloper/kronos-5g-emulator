package nasType

// IEExtendedProtocolConfigurationOptions identifies the optional ExtendedProtocolConfigurationOptions IE on the wire.
const IEExtendedProtocolConfigurationOptions uint8 = 0x7b

// ExtendedProtocolConfigurationOptions is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type ExtendedProtocolConfigurationOptions struct {
	Value []byte
}
