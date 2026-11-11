// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LoggedMDTNrExtIEsEntryPresent uint32

const (
	LoggedMDTNrExtIEsEntryPresentNothing LoggedMDTNrExtIEsEntryPresent = iota
	LoggedMDTNrExtIEsEntryPresentEarlyMeasurement
	LoggedMDTNrExtIEsEntryPresentUnknown
)

type LoggedMDTNrExtIEsEntry struct {
	ID               ProtocolExtensionID
	Criticality      Criticality
	Present          LoggedMDTNrExtIEsEntryPresent
	EarlyMeasurement EarlyMeasurement `ngap:"id:360,criticality:ignore,presence:optional"`
	UnknownValue     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LoggedMDTNrExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LoggedMDTNr-ExtIEsEntry", v)
}

// DecodeLoggedMDTNrExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLoggedMDTNrExtIEsEntry(data []byte) (LoggedMDTNrExtIEsEntry, int, error) {
	var v LoggedMDTNrExtIEsEntry
	err := ngapcodec.DecodeParameter("LoggedMDTNr-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LoggedMDTNrExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
