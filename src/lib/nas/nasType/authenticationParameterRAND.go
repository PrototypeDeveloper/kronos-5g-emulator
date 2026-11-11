package nasType

// IEAuthenticationParameterRAND identifies the optional AuthenticationParameterRAND IE on the wire.
const IEAuthenticationParameterRAND uint8 = 0x21

// AuthenticationParameterRAND is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type AuthenticationParameterRAND struct {
	Value []byte
}
