// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type InterSystemFailureIndicationExtIEsEntryPresent uint32

const (
	InterSystemFailureIndicationExtIEsEntryPresentNothing InterSystemFailureIndicationExtIEsEntryPresent = iota
	InterSystemFailureIndicationExtIEsEntryPresentUnknown
)

type InterSystemFailureIndicationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      InterSystemFailureIndicationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InterSystemFailureIndicationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InterSystemFailureIndication-ExtIEsEntry", v)
}

// DecodeInterSystemFailureIndicationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeInterSystemFailureIndicationExtIEsEntry(data []byte) (InterSystemFailureIndicationExtIEsEntry, int, error) {
	var v InterSystemFailureIndicationExtIEsEntry
	err := ngapcodec.DecodeParameter("InterSystemFailureIndication-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return InterSystemFailureIndicationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
