// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaOfInterestExtIEsEntryPresent uint32

const (
	AreaOfInterestExtIEsEntryPresentNothing AreaOfInterestExtIEsEntryPresent = iota
	AreaOfInterestExtIEsEntryPresentUnknown
)

type AreaOfInterestExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AreaOfInterestExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterest-ExtIEsEntry", v)
}

// DecodeAreaOfInterestExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestExtIEsEntry(data []byte) (AreaOfInterestExtIEsEntry, int, error) {
	var v AreaOfInterestExtIEsEntry
	err := ngapcodec.DecodeParameter("AreaOfInterest-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
