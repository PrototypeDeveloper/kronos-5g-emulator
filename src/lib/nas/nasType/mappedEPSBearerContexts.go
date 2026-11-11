package nasType

// IEMappedEPSBearerContexts identifies the optional MappedEPSBearerContexts IE on the wire.
const IEMappedEPSBearerContexts uint8 = 0x75

// MappedEPSBearerContexts is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type MappedEPSBearerContexts struct {
	Value []byte
}
