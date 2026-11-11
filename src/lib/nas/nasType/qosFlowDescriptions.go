package nasType

// IEQoSFlowDescriptions identifies the Authorized QoS flow descriptions IE.
const IEQoSFlowDescriptions uint8 = 0x79

// QoSFlowDescriptions contains value bytes without the IEI or length.
// Presence is controlled by the containing message's Has flag.
type QoSFlowDescriptions struct {
	Value []byte
}
