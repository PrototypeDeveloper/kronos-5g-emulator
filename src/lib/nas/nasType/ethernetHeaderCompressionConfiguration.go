package nasType

// IEEthernetHeaderCompressionConfiguration identifies the optional EthernetHeaderCompressionConfiguration IE on the wire.
const IEEthernetHeaderCompressionConfiguration uint8 = 0x1f

// EthernetHeaderCompressionConfiguration is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type EthernetHeaderCompressionConfiguration struct {
	Value uint8
}
