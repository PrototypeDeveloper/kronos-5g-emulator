// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AMFPagingTargetExtIEsEntryPresent uint32

const (
	AMFPagingTargetExtIEsEntryPresentNothing AMFPagingTargetExtIEsEntryPresent = iota
	AMFPagingTargetExtIEsEntryPresentUnknown
)

type AMFPagingTargetExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      AMFPagingTargetExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFPagingTargetExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMFPagingTarget-ExtIEsEntry", v)
}

// DecodeAMFPagingTargetExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAMFPagingTargetExtIEsEntry(data []byte) (AMFPagingTargetExtIEsEntry, int, error) {
	var v AMFPagingTargetExtIEsEntry
	err := ngapcodec.DecodeParameter("AMFPagingTarget-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFPagingTargetExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
