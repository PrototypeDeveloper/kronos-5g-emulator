// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PagingPolicyDifferentiationExtIEsEntryPresent uint32

const (
	PagingPolicyDifferentiationExtIEsEntryPresentNothing PagingPolicyDifferentiationExtIEsEntryPresent = iota
	PagingPolicyDifferentiationExtIEsEntryPresentUnknown
)

type PagingPolicyDifferentiationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PagingPolicyDifferentiationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingPolicyDifferentiationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingPolicyDifferentiation-ExtIEsEntry", v)
}

// DecodePagingPolicyDifferentiationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePagingPolicyDifferentiationExtIEsEntry(data []byte) (PagingPolicyDifferentiationExtIEsEntry, int, error) {
	var v PagingPolicyDifferentiationExtIEsEntry
	err := ngapcodec.DecodeParameter("PagingPolicyDifferentiation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingPolicyDifferentiationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
