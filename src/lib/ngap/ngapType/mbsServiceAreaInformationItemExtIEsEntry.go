// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSServiceAreaInformationItemExtIEsEntryPresent uint32

const (
	MBSServiceAreaInformationItemExtIEsEntryPresentNothing MBSServiceAreaInformationItemExtIEsEntryPresent = iota
	MBSServiceAreaInformationItemExtIEsEntryPresentUnknown
)

type MBSServiceAreaInformationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSServiceAreaInformationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSServiceAreaInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ServiceAreaInformationItem-ExtIEsEntry", v)
}

// DecodeMBSServiceAreaInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSServiceAreaInformationItemExtIEsEntry(data []byte) (MBSServiceAreaInformationItemExtIEsEntry, int, error) {
	var v MBSServiceAreaInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-ServiceAreaInformationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSServiceAreaInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
