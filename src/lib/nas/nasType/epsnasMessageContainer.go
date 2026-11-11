package nasType

// IEEPSNASMessageContainer identifies the optional EPSNASMessageContainer IE on the wire.
const IEEPSNASMessageContainer uint8 = 0x70

// EPSNASMessageContainer is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type EPSNASMessageContainer struct {
	Value []byte
}
