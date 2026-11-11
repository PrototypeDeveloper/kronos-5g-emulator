package nasType

// IEAdditional5GSecurityInformation identifies the optional Additional5GSecurityInformation IE on the wire.
const IEAdditional5GSecurityInformation uint8 = 0x36

// Additional5GSecurityInformation is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type Additional5GSecurityInformation struct {
	Value uint8
}
