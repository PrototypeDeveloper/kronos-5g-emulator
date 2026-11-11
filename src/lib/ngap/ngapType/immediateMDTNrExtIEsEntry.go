// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ImmediateMDTNrExtIEsEntryPresent uint32

const (
	ImmediateMDTNrExtIEsEntryPresentNothing ImmediateMDTNrExtIEsEntryPresent = iota
	ImmediateMDTNrExtIEsEntryPresentUnknown
)

type ImmediateMDTNrExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ImmediateMDTNrExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ImmediateMDTNrExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ImmediateMDTNr-ExtIEsEntry", v)
}

// DecodeImmediateMDTNrExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeImmediateMDTNrExtIEsEntry(data []byte) (ImmediateMDTNrExtIEsEntry, int, error) {
	var v ImmediateMDTNrExtIEsEntry
	err := ngapcodec.DecodeParameter("ImmediateMDTNr-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ImmediateMDTNrExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
