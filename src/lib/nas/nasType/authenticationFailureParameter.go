package nasType

// IEAuthenticationFailureParameter identifies the optional AuthenticationFailureParameter IE on the wire.
const IEAuthenticationFailureParameter uint8 = 0x30

// AuthenticationFailureParameter is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type AuthenticationFailureParameter struct {
	Value []byte
}
