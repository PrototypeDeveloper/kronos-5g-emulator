package nasType

// IESelectedEPSNASAlgorithms identifies the optional SelectedEPSNASAlgorithms IE on the wire.
const IESelectedEPSNASAlgorithms uint8 = 0x57

// SelectedEPSNASAlgorithms is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type SelectedEPSNASAlgorithms struct {
	Value uint8
}
