// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingPolicyDifferentiationItem represents ASN.1 PagingPolicyDifferentiationItem.
type PagingPolicyDifferentiationItem struct {
	HasQosFlowIdentifier              bool
	QosFlowIdentifier                 QosFlowIdentifier
	HasPagingPolicyIndicator          bool
	PagingPolicyIndicator             PagingPolicyIndicator
	HasAllocationAndRetentionPriority bool
	AllocationAndRetentionPriority    AllocationAndRetentionPriority
	HasFiveQI                         bool
	FiveQI                            FiveQI
	HasDlDataSize                     bool
	DlDataSize                        int64 `aper:"valueExt,valueLB:0,valueUB:96000"`
	HasIEExtensions                   bool
	IEExtensions                      []PagingPolicyDifferentiationItemExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingPolicyDifferentiationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingPolicyDifferentiationItem", v)
}

// DecodePagingPolicyDifferentiationItem decodes a standalone APER value and reports the consumed octets.
func DecodePagingPolicyDifferentiationItem(data []byte) (PagingPolicyDifferentiationItem, int, error) {
	var v PagingPolicyDifferentiationItem
	err := ngapcodec.DecodeParameter("PagingPolicyDifferentiationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingPolicyDifferentiationItem{}, 0, err
	}
	return v, len(data), err
}
