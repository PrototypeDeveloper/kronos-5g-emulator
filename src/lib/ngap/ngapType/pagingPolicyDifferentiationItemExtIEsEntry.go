// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PagingPolicyDifferentiationItemExtIEsEntryPresent uint32

const (
	PagingPolicyDifferentiationItemExtIEsEntryPresentNothing PagingPolicyDifferentiationItemExtIEsEntryPresent = iota
	PagingPolicyDifferentiationItemExtIEsEntryPresentUnknown
)

type PagingPolicyDifferentiationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PagingPolicyDifferentiationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingPolicyDifferentiationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingPolicyDifferentiationItem-ExtIEsEntry", v)
}

// DecodePagingPolicyDifferentiationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePagingPolicyDifferentiationItemExtIEsEntry(data []byte) (PagingPolicyDifferentiationItemExtIEsEntry, int, error) {
	var v PagingPolicyDifferentiationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PagingPolicyDifferentiationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingPolicyDifferentiationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
