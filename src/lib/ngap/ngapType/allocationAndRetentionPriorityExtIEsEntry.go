// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AllocationAndRetentionPriorityExtIEsEntryPresent uint32

const (
	AllocationAndRetentionPriorityExtIEsEntryPresentNothing AllocationAndRetentionPriorityExtIEsEntryPresent = iota
	AllocationAndRetentionPriorityExtIEsEntryPresentUnknown
)

type AllocationAndRetentionPriorityExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AllocationAndRetentionPriorityExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllocationAndRetentionPriorityExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AllocationAndRetentionPriority-ExtIEsEntry", v)
}

// DecodeAllocationAndRetentionPriorityExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAllocationAndRetentionPriorityExtIEsEntry(data []byte) (AllocationAndRetentionPriorityExtIEsEntry, int, error) {
	var v AllocationAndRetentionPriorityExtIEsEntry
	err := ngapcodec.DecodeParameter("AllocationAndRetentionPriority-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllocationAndRetentionPriorityExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
