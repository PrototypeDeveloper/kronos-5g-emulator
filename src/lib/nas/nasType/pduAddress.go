package nasType

// IESuggestedInterfaceIdentifier identifies the optional SuggestedInterfaceIdentifier IE on the wire.
const IESuggestedInterfaceIdentifier uint8 = 0x29

// IEPDUAddress identifies the optional PDUAddress IE on the wire.
const IEPDUAddress uint8 = 0x29

// PDUAddress is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type PDUAddress struct {
	Value []byte
}
