// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllocationAndRetentionPriority represents ASN.1 AllocationAndRetentionPriority.
type AllocationAndRetentionPriority struct {
	PriorityLevelARP        PriorityLevelARP
	PreEmptionCapability    PreEmptionCapability
	PreEmptionVulnerability PreEmptionVulnerability
	HasIEExtensions         bool
	IEExtensions            []AllocationAndRetentionPriorityExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllocationAndRetentionPriority) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AllocationAndRetentionPriority", v)
}

// DecodeAllocationAndRetentionPriority decodes a standalone APER value and reports the consumed octets.
func DecodeAllocationAndRetentionPriority(data []byte) (AllocationAndRetentionPriority, int, error) {
	var v AllocationAndRetentionPriority
	err := ngapcodec.DecodeParameter("AllocationAndRetentionPriority", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllocationAndRetentionPriority{}, 0, err
	}
	return v, len(data), err
}
