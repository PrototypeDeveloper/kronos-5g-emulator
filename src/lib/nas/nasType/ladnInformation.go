package nasType

// IELADNInformation identifies the optional LADNInformation IE on the wire.
const IELADNInformation uint8 = 0x79

// LADNInformation is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type LADNInformation struct {
	Value []byte
}
