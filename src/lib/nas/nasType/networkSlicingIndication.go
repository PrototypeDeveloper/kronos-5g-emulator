package nasType

// IENetworkSlicingIndication identifies the optional NetworkSlicingIndication IE on the wire.
const IENetworkSlicingIndication uint8 = 0x90

// NetworkSlicingIndication is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type NetworkSlicingIndication struct {
	Value uint8
}
