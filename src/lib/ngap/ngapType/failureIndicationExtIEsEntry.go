// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FailureIndicationExtIEsEntryPresent uint32

const (
	FailureIndicationExtIEsEntryPresentNothing FailureIndicationExtIEsEntryPresent = iota
	FailureIndicationExtIEsEntryPresentUnknown
)

type FailureIndicationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      FailureIndicationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FailureIndicationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FailureIndication-ExtIEsEntry", v)
}

// DecodeFailureIndicationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFailureIndicationExtIEsEntry(data []byte) (FailureIndicationExtIEsEntry, int, error) {
	var v FailureIndicationExtIEsEntry
	err := ngapcodec.DecodeParameter("FailureIndication-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FailureIndicationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
