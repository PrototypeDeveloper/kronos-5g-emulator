// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type M7ConfigurationExtIEsEntryPresent uint32

const (
	M7ConfigurationExtIEsEntryPresentNothing M7ConfigurationExtIEsEntryPresent = iota
	M7ConfigurationExtIEsEntryPresentM7ReportAmount
	M7ConfigurationExtIEsEntryPresentUnknown
)

type M7ConfigurationExtIEsEntry struct {
	ID             ProtocolExtensionID
	Criticality    Criticality
	Present        M7ConfigurationExtIEsEntryPresent
	M7ReportAmount M7ReportAmountMDT `ngap:"id:339,criticality:ignore,presence:optional"`
	UnknownValue   []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M7ConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M7Configuration-ExtIEsEntry", v)
}

// DecodeM7ConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeM7ConfigurationExtIEsEntry(data []byte) (M7ConfigurationExtIEsEntry, int, error) {
	var v M7ConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("M7Configuration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return M7ConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
