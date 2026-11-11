package nasType

// IEPDUSessionID identifies the optional PDUSessionID IE on the wire.
const IEPDUSessionID uint8 = 0x12

// IEOldPDUSessionID identifies the optional OldPDUSessionID IE on the wire.
const IEOldPDUSessionID uint8 = 0x59

// PDUSessionID is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type PDUSessionID struct {
	Value uint8
}
