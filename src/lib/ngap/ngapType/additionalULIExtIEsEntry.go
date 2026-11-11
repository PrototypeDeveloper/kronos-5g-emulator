// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AdditionalULIExtIEsEntryPresent uint32

const (
	AdditionalULIExtIEsEntryPresentNothing AdditionalULIExtIEsEntryPresent = iota
	AdditionalULIExtIEsEntryPresentUnknown
)

type AdditionalULIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AdditionalULIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalULIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AdditionalULI-ExtIEsEntry", v)
}

// DecodeAdditionalULIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalULIExtIEsEntry(data []byte) (AdditionalULIExtIEsEntry, int, error) {
	var v AdditionalULIExtIEsEntry
	err := ngapcodec.DecodeParameter("AdditionalULI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalULIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
