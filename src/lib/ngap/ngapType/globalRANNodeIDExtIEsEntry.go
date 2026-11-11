// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalRANNodeIDExtIEsEntryPresent uint32

const (
	GlobalRANNodeIDExtIEsEntryPresentNothing GlobalRANNodeIDExtIEsEntryPresent = iota
	GlobalRANNodeIDExtIEsEntryPresentGlobalTNGFID
	GlobalRANNodeIDExtIEsEntryPresentGlobalTWIFID
	GlobalRANNodeIDExtIEsEntryPresentGlobalWAGFID
	GlobalRANNodeIDExtIEsEntryPresentUnknown
)

type GlobalRANNodeIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      GlobalRANNodeIDExtIEsEntryPresent
	GlobalTNGFID GlobalTNGFID `ngap:"id:240,criticality:reject,presence:mandatory"`
	GlobalTWIFID GlobalTWIFID `ngap:"id:241,criticality:reject,presence:mandatory"`
	GlobalWAGFID GlobalWAGFID `ngap:"id:242,criticality:reject,presence:mandatory"`
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalRANNodeIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalRANNodeID-ExtIEsEntry", v)
}

// DecodeGlobalRANNodeIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalRANNodeIDExtIEsEntry(data []byte) (GlobalRANNodeIDExtIEsEntry, int, error) {
	var v GlobalRANNodeIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalRANNodeID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalRANNodeIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
