// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DAPSResponseInfoExtIEsEntryPresent uint32

const (
	DAPSResponseInfoExtIEsEntryPresentNothing DAPSResponseInfoExtIEsEntryPresent = iota
	DAPSResponseInfoExtIEsEntryPresentUnknown
)

type DAPSResponseInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DAPSResponseInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DAPSResponseInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DAPSResponseInfo-ExtIEsEntry", v)
}

// DecodeDAPSResponseInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDAPSResponseInfoExtIEsEntry(data []byte) (DAPSResponseInfoExtIEsEntry, int, error) {
	var v DAPSResponseInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("DAPSResponseInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSResponseInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
