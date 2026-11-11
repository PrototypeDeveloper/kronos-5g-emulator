package nasType

// IEPDUSessionReactivationResultErrorCause identifies the optional PDUSessionReactivationResultErrorCause IE on the wire.
const IEPDUSessionReactivationResultErrorCause uint8 = 0x72

// PDUSessionReactivationResultErrorCause is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type PDUSessionReactivationResultErrorCause struct {
	Value []byte
}
