package nasType

// IEPortManagementInformationContainer identifies the optional PortManagementInformationContainer IE on the wire.
const IEPortManagementInformationContainer uint8 = 0x74

// PortManagementInformationContainer is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type PortManagementInformationContainer struct {
	Value []byte
}
