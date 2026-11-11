package nasType

// IEOperatorDefinedAccessCategoryDefinitions identifies the optional OperatorDefinedAccessCategoryDefinitions IE on the wire.
const IEOperatorDefinedAccessCategoryDefinitions uint8 = 0x76

// OperatorDefinedAccessCategoryDefinitions is a named NAS information element.
// Value owns the IE value bytes, excluding IEI and length; internal subfields remain opaque.
// Presence is controlled by the containing message's corresponding Has flag.
type OperatorDefinedAccessCategoryDefinitions struct {
	Value []byte
}
