package nasType

// IELastVisitedRegisteredTAI identifies the optional LastVisitedRegisteredTAI IE on the wire.
const IELastVisitedRegisteredTAI uint8 = 0x52

// LastVisitedRegisteredTAI is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type LastVisitedRegisteredTAI struct {
	Value []byte
}
