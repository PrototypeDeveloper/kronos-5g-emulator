package nasType

// IENetworkFeatureSupport5GS identifies the optional NetworkFeatureSupport5GS IE on the wire.
const IENetworkFeatureSupport5GS uint8 = 0x21

// NetworkFeatureSupport5GS is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type NetworkFeatureSupport5GS struct {
	Value []byte
}
