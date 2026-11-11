package nasType

// IESMPDUDNRequestContainer identifies the optional SMPDUDNRequestContainer IE on the wire.
const IESMPDUDNRequestContainer uint8 = 0x39

// SMPDUDNRequestContainer is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type SMPDUDNRequestContainer struct {
	Value []byte
}
