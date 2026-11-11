package nasType

// IEAlwaysOnPDUSessionRequested identifies the optional AlwaysOnPDUSessionRequested IE on the wire.
const IEAlwaysOnPDUSessionRequested uint8 = 0xb0

// AlwaysOnPDUSessionRequested is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type AlwaysOnPDUSessionRequested struct {
	Value uint8
}
