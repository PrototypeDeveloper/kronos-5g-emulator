// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserLocationInformationWAGFExtIEsEntryPresent uint32

const (
	UserLocationInformationWAGFExtIEsEntryPresentNothing UserLocationInformationWAGFExtIEsEntryPresent = iota
	UserLocationInformationWAGFExtIEsEntryPresentGlobalCableID
	UserLocationInformationWAGFExtIEsEntryPresentHFCNodeIDNew
	UserLocationInformationWAGFExtIEsEntryPresentGlobalCableIDNew
	UserLocationInformationWAGFExtIEsEntryPresentUnknown
)

type UserLocationInformationWAGFExtIEsEntry struct {
	ID               ProtocolIEID
	Criticality      Criticality
	Present          UserLocationInformationWAGFExtIEsEntryPresent
	GlobalCableID    GlobalCableID    `ngap:"id:275,criticality:ignore,presence:mandatory"`
	HFCNodeIDNew     HFCNodeIDNew     `ngap:"id:362,criticality:ignore,presence:mandatory"`
	GlobalCableIDNew GlobalCableIDNew `ngap:"id:363,criticality:ignore,presence:mandatory"`
	UnknownValue     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationWAGFExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationW-AGF-ExtIEsEntry", v)
}

// DecodeUserLocationInformationWAGFExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationWAGFExtIEsEntry(data []byte) (UserLocationInformationWAGFExtIEsEntry, int, error) {
	var v UserLocationInformationWAGFExtIEsEntry
	err := ngapcodec.DecodeParameter("UserLocationInformationW-AGF-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationWAGFExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
