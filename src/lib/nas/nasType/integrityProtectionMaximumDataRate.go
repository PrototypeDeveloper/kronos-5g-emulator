package nasType

// IEIntegrityProtectionMaximumDataRate identifies the optional IntegrityProtectionMaximumDataRate IE on the wire.
const IEIntegrityProtectionMaximumDataRate uint8 = 0x13

// IntegrityProtectionMaximumDataRate is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type IntegrityProtectionMaximumDataRate struct {
	Value []byte
}
