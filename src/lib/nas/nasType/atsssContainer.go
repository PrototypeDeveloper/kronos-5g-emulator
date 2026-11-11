package nasType

// IEATSSSContainer identifies the optional ATSSSContainer IE on the wire.
const IEATSSSContainer uint8 = 0x77

// ATSSSContainer is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type ATSSSContainer struct {
	Value []byte
}
