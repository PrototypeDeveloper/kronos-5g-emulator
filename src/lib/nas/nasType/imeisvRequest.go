package nasType

// IEIMEISVRequest identifies the optional IMEISVRequest IE on the wire.
const IEIMEISVRequest uint8 = 0xe0

// IMEISVRequest is a named NAS information element.
// Value holds the half-octet or octet value; the message codec supplies IEI and length.
// Presence is controlled by the containing message's corresponding Has flag.
type IMEISVRequest struct {
	Value uint8
}
