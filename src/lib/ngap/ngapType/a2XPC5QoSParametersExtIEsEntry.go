// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type A2XPC5QoSParametersExtIEsEntryPresent uint32

const (
	A2XPC5QoSParametersExtIEsEntryPresentNothing A2XPC5QoSParametersExtIEsEntryPresent = iota
	A2XPC5QoSParametersExtIEsEntryPresentUnknown
)

type A2XPC5QoSParametersExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      A2XPC5QoSParametersExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v A2XPC5QoSParametersExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("A2X-PC5-QoS-Parameters-ExtIEsEntry", v)
}

// DecodeA2XPC5QoSParametersExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeA2XPC5QoSParametersExtIEsEntry(data []byte) (A2XPC5QoSParametersExtIEsEntry, int, error) {
	var v A2XPC5QoSParametersExtIEsEntry
	err := ngapcodec.DecodeParameter("A2X-PC5-QoS-Parameters-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return A2XPC5QoSParametersExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
