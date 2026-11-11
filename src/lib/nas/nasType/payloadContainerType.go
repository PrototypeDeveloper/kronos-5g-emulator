package nasType

// IEPayloadContainerType identifies the optional PayloadContainerType IE on the wire.
const IEPayloadContainerType uint8 = 0x80

// PayloadContainerType is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type PayloadContainerType struct {
	Value uint8
}
