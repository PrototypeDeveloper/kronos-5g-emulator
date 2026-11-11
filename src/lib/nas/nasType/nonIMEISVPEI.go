package nasType

// IENonIMEISVPEI identifies the optional NonIMEISVPEI IE on the wire.
const IENonIMEISVPEI uint8 = 0x78

// NonIMEISVPEI is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type NonIMEISVPEI struct {
	Value []byte
}
