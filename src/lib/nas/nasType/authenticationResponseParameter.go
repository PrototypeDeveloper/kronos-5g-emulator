package nasType

// IEAuthenticationResponseParameter identifies the optional AuthenticationResponseParameter IE on the wire.
const IEAuthenticationResponseParameter uint8 = 0x2d

// AuthenticationResponseParameter is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type AuthenticationResponseParameter struct {
	Value []byte
}
