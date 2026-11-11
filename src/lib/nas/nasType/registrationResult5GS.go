package nasType

// IERegistrationResult5GS identifies the optional RegistrationResult5GS IE on the wire.
const IERegistrationResult5GS uint8 = 0x44

// RegistrationResult5GS is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type RegistrationResult5GS struct {
	Value uint8
}
