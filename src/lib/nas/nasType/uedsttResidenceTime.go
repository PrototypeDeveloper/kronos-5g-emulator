package nasType

// IEUEDSTTResidenceTime identifies the optional UEDSTTResidenceTime IE on the wire.
const IEUEDSTTResidenceTime uint8 = 0x6f

// UEDSTTResidenceTime is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type UEDSTTResidenceTime struct {
	Value []byte
}
