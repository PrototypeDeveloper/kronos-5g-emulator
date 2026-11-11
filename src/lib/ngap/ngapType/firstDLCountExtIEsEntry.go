// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FirstDLCountExtIEsEntryPresent uint32

const (
	FirstDLCountExtIEsEntryPresentNothing FirstDLCountExtIEsEntryPresent = iota
	FirstDLCountExtIEsEntryPresentUnknown
)

type FirstDLCountExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      FirstDLCountExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FirstDLCountExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FirstDLCount-ExtIEsEntry", v)
}

// DecodeFirstDLCountExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFirstDLCountExtIEsEntry(data []byte) (FirstDLCountExtIEsEntry, int, error) {
	var v FirstDLCountExtIEsEntry
	err := ngapcodec.DecodeParameter("FirstDLCount-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FirstDLCountExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
