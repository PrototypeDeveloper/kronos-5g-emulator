// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingPriority represents ASN.1 PagingPriority.
const (
	PagingPriorityPresentPriolevel1 Enumerated = 0
	PagingPriorityPresentPriolevel2 Enumerated = 1
	PagingPriorityPresentPriolevel3 Enumerated = 2
	PagingPriorityPresentPriolevel4 Enumerated = 3
	PagingPriorityPresentPriolevel5 Enumerated = 4
	PagingPriorityPresentPriolevel6 Enumerated = 5
	PagingPriorityPresentPriolevel7 Enumerated = 6
	PagingPriorityPresentPriolevel8 Enumerated = 7
)

type PagingPriority struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:7"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingPriority) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingPriority", v)
}

// DecodePagingPriority decodes a standalone APER value and reports the consumed octets.
func DecodePagingPriority(data []byte) (PagingPriority, int, error) {
	var v PagingPriority
	err := ngapcodec.DecodeParameter("PagingPriority", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingPriority{}, 0, err
	}
	return v, len(data), err
}
