// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LoggedMDTTriggerExtIEsEntryPresent uint32

const (
	LoggedMDTTriggerExtIEsEntryPresentNothing LoggedMDTTriggerExtIEsEntryPresent = iota
	LoggedMDTTriggerExtIEsEntryPresentUnknown
)

type LoggedMDTTriggerExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      LoggedMDTTriggerExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LoggedMDTTriggerExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LoggedMDTTrigger-ExtIEsEntry", v)
}

// DecodeLoggedMDTTriggerExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLoggedMDTTriggerExtIEsEntry(data []byte) (LoggedMDTTriggerExtIEsEntry, int, error) {
	var v LoggedMDTTriggerExtIEsEntry
	err := ngapcodec.DecodeParameter("LoggedMDTTrigger-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LoggedMDTTriggerExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
