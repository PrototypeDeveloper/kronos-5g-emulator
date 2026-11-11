package nasType

// IERejectedNSSAI identifies the optional RejectedNSSAI IE on the wire.
const IERejectedNSSAI uint8 = 0x11

// IEDeregistrationRejectedNSSAI identifies the Rejected NSSAI IE of DeregistrationRequestNW on the wire.
const IEDeregistrationRejectedNSSAI uint8 = 0x6d

// RejectedNSSAI is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type RejectedNSSAI struct {
	Value []byte
}
