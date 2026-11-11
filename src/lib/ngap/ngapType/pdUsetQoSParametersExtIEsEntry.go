// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUsetQoSParametersExtIEsEntryPresent uint32

const (
	PDUsetQoSParametersExtIEsEntryPresentNothing PDUsetQoSParametersExtIEsEntryPresent = iota
	PDUsetQoSParametersExtIEsEntryPresentUnknown
)

type PDUsetQoSParametersExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUsetQoSParametersExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUsetQoSParametersExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUsetQoSParameters-ExtIEsEntry", v)
}

// DecodePDUsetQoSParametersExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUsetQoSParametersExtIEsEntry(data []byte) (PDUsetQoSParametersExtIEsEntry, int, error) {
	var v PDUsetQoSParametersExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUsetQoSParameters-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUsetQoSParametersExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
