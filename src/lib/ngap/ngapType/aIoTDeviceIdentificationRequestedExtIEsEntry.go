// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTDeviceIdentificationRequestedExtIEsEntryPresent uint32

const (
	AIoTDeviceIdentificationRequestedExtIEsEntryPresentNothing AIoTDeviceIdentificationRequestedExtIEsEntryPresent = iota
	AIoTDeviceIdentificationRequestedExtIEsEntryPresentUnknown
)

type AIoTDeviceIdentificationRequestedExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      AIoTDeviceIdentificationRequestedExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTDeviceIdentificationRequestedExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-DeviceIdentificationRequested-ExtIEsEntry", v)
}

// DecodeAIoTDeviceIdentificationRequestedExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTDeviceIdentificationRequestedExtIEsEntry(data []byte) (AIoTDeviceIdentificationRequestedExtIEsEntry, int, error) {
	var v AIoTDeviceIdentificationRequestedExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoT-DeviceIdentificationRequested-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTDeviceIdentificationRequestedExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
