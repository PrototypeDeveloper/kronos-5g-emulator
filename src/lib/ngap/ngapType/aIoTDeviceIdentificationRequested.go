// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTDeviceIdentificationRequested represents ASN.1 AIoT-DeviceIdentificationRequested.
type AIoTDeviceIdentificationRequestedPresent uint32

const (
	AIoTDeviceIdentificationRequestedPresentNothing AIoTDeviceIdentificationRequestedPresent = iota
	AIoTDeviceIdentificationRequestedPresentSingleDevice
	AIoTDeviceIdentificationRequestedPresentGroupDevices
	AIoTDeviceIdentificationRequestedPresentAllDevices
	AIoTDeviceIdentificationRequestedPresentChoiceExtensions
)

type AIoTDeviceIdentificationRequested struct {
	Present          AIoTDeviceIdentificationRequestedPresent
	SingleDevice     BitString `aper:"sizeExt,sizeLB:1,sizeUB:1024"`
	GroupDevices     BitString `aper:"sizeExt,sizeLB:1,sizeUB:1024"`
	AllDevices       struct{}
	ChoiceExtensions AIoTDeviceIdentificationRequestedExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTDeviceIdentificationRequested) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-DeviceIdentificationRequested", v)
}

// DecodeAIoTDeviceIdentificationRequested decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTDeviceIdentificationRequested(data []byte) (AIoTDeviceIdentificationRequested, int, error) {
	var v AIoTDeviceIdentificationRequested
	err := ngapcodec.DecodeParameter("AIoT-DeviceIdentificationRequested", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTDeviceIdentificationRequested{}, 0, err
	}
	return v, len(data), err
}
