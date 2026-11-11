// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTDeviceReportItemExtIEsEntryPresent uint32

const (
	AIoTDeviceReportItemExtIEsEntryPresentNothing AIoTDeviceReportItemExtIEsEntryPresent = iota
	AIoTDeviceReportItemExtIEsEntryPresentUnknown
)

type AIoTDeviceReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIoTDeviceReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTDeviceReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-DeviceReportItem-ExtIEsEntry", v)
}

// DecodeAIoTDeviceReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTDeviceReportItemExtIEsEntry(data []byte) (AIoTDeviceReportItemExtIEsEntry, int, error) {
	var v AIoTDeviceReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoT-DeviceReportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTDeviceReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
