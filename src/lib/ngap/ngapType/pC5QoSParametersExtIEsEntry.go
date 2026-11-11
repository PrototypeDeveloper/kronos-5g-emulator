// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PC5QoSParametersExtIEsEntryPresent uint32

const (
	PC5QoSParametersExtIEsEntryPresentNothing PC5QoSParametersExtIEsEntryPresent = iota
	PC5QoSParametersExtIEsEntryPresentUnknown
)

type PC5QoSParametersExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PC5QoSParametersExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PC5QoSParametersExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PC5QoSParameters-ExtIEsEntry", v)
}

// DecodePC5QoSParametersExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePC5QoSParametersExtIEsEntry(data []byte) (PC5QoSParametersExtIEsEntry, int, error) {
	var v PC5QoSParametersExtIEsEntry
	err := ngapcodec.DecodeParameter("PC5QoSParameters-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PC5QoSParametersExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
