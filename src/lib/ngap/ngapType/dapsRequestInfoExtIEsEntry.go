// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DAPSRequestInfoExtIEsEntryPresent uint32

const (
	DAPSRequestInfoExtIEsEntryPresentNothing DAPSRequestInfoExtIEsEntryPresent = iota
	DAPSRequestInfoExtIEsEntryPresentUnknown
)

type DAPSRequestInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DAPSRequestInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DAPSRequestInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DAPSRequestInfo-ExtIEsEntry", v)
}

// DecodeDAPSRequestInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDAPSRequestInfoExtIEsEntry(data []byte) (DAPSRequestInfoExtIEsEntry, int, error) {
	var v DAPSRequestInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("DAPSRequestInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSRequestInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
