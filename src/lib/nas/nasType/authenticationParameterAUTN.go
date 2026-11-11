package nasType

// IEAuthenticationParameterAUTN identifies the optional AuthenticationParameterAUTN IE on the wire.
const IEAuthenticationParameterAUTN uint8 = 0x20

// AuthenticationParameterAUTN is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type AuthenticationParameterAUTN struct {
	Value []byte
}
