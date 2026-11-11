// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaOfInterestRANNodeItemExtIEsEntryPresent uint32

const (
	AreaOfInterestRANNodeItemExtIEsEntryPresentNothing AreaOfInterestRANNodeItemExtIEsEntryPresent = iota
	AreaOfInterestRANNodeItemExtIEsEntryPresentUnknown
)

type AreaOfInterestRANNodeItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AreaOfInterestRANNodeItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestRANNodeItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestRANNodeItem-ExtIEsEntry", v)
}

// DecodeAreaOfInterestRANNodeItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestRANNodeItemExtIEsEntry(data []byte) (AreaOfInterestRANNodeItemExtIEsEntry, int, error) {
	var v AreaOfInterestRANNodeItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AreaOfInterestRANNodeItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestRANNodeItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
