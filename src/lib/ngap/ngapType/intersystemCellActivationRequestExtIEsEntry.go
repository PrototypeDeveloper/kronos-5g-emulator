// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemCellActivationRequestExtIEsEntryPresent uint32

const (
	IntersystemCellActivationRequestExtIEsEntryPresentNothing IntersystemCellActivationRequestExtIEsEntryPresent = iota
	IntersystemCellActivationRequestExtIEsEntryPresentUnknown
)

type IntersystemCellActivationRequestExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemCellActivationRequestExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemCellActivationRequestExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemCellActivationRequest-ExtIEsEntry", v)
}

// DecodeIntersystemCellActivationRequestExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemCellActivationRequestExtIEsEntry(data []byte) (IntersystemCellActivationRequestExtIEsEntry, int, error) {
	var v IntersystemCellActivationRequestExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemCellActivationRequest-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemCellActivationRequestExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
