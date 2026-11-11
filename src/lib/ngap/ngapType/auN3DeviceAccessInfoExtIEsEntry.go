// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AUN3DeviceAccessInfoExtIEsEntryPresent uint32

const (
	AUN3DeviceAccessInfoExtIEsEntryPresentNothing AUN3DeviceAccessInfoExtIEsEntryPresent = iota
	AUN3DeviceAccessInfoExtIEsEntryPresentUnknown
)

type AUN3DeviceAccessInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AUN3DeviceAccessInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AUN3DeviceAccessInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AUN3DeviceAccessInfo-ExtIEsEntry", v)
}

// DecodeAUN3DeviceAccessInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAUN3DeviceAccessInfoExtIEsEntry(data []byte) (AUN3DeviceAccessInfoExtIEsEntry, int, error) {
	var v AUN3DeviceAccessInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("AUN3DeviceAccessInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AUN3DeviceAccessInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
