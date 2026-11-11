package nasType

// IEPDUSessionStatus identifies the optional PDUSessionStatus IE on the wire.
const IEPDUSessionStatus uint8 = 0x50

// PDUSessionStatus is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type PDUSessionStatus struct {
	Value []byte
}
