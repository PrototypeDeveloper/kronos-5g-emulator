package nasType

// IEPDUSessionType identifies the optional PDUSessionType IE on the wire.
const IEPDUSessionType uint8 = 0x90

// PDUSessionType is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type PDUSessionType struct {
	Value uint8
}
