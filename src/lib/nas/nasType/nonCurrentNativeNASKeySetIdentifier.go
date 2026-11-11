package nasType

// IENonCurrentNativeNASKeySetIdentifier identifies the optional NonCurrentNativeNASKeySetIdentifier IE on the wire.
const IENonCurrentNativeNASKeySetIdentifier uint8 = 0xc0

// NonCurrentNativeNASKeySetIdentifier is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type NonCurrentNativeNASKeySetIdentifier struct {
	Value uint8
}
