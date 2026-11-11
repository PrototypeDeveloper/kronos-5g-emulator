package nasType

// IEPDUSessionReactivationResult identifies the optional PDUSessionReactivationResult IE on the wire.
const IEPDUSessionReactivationResult uint8 = 0x26

// PDUSessionReactivationResult is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type PDUSessionReactivationResult struct {
	Value []byte
}
