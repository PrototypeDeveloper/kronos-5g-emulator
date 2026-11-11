package nasType

// IEDSTTEthernetPortMACAddress identifies the optional DSTTEthernetPortMACAddress IE on the wire.
const IEDSTTEthernetPortMACAddress uint8 = 0x6e

// DSTTEthernetPortMACAddress is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type DSTTEthernetPortMACAddress struct {
	Value []byte
}
