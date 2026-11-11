// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionSetupResponseItemExtIEsEntryPresent uint32

const (
	MBSSessionSetupResponseItemExtIEsEntryPresentNothing MBSSessionSetupResponseItemExtIEsEntryPresent = iota
	MBSSessionSetupResponseItemExtIEsEntryPresentUnknown
)

type MBSSessionSetupResponseItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSSessionSetupResponseItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupResponseItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupResponseItem-ExtIEsEntry", v)
}

// DecodeMBSSessionSetupResponseItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupResponseItemExtIEsEntry(data []byte) (MBSSessionSetupResponseItemExtIEsEntry, int, error) {
	var v MBSSessionSetupResponseItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionSetupResponseItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupResponseItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
