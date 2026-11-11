package nasType

// IELADNIndication identifies the optional LADNIndication IE on the wire.
const IELADNIndication uint8 = 0x74

// LADNIndication is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type LADNIndication struct {
	Value []byte
}
