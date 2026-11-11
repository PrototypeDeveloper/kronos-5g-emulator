// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type M5ConfigurationExtIEsEntryPresent uint32

const (
	M5ConfigurationExtIEsEntryPresentNothing M5ConfigurationExtIEsEntryPresent = iota
	M5ConfigurationExtIEsEntryPresentM5ReportAmount
	M5ConfigurationExtIEsEntryPresentUnknown
)

type M5ConfigurationExtIEsEntry struct {
	ID             ProtocolExtensionID
	Criticality    Criticality
	Present        M5ConfigurationExtIEsEntryPresent
	M5ReportAmount M5ReportAmountMDT `ngap:"id:337,criticality:ignore,presence:optional"`
	UnknownValue   []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M5ConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M5Configuration-ExtIEsEntry", v)
}

// DecodeM5ConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeM5ConfigurationExtIEsEntry(data []byte) (M5ConfigurationExtIEsEntry, int, error) {
	var v M5ConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("M5Configuration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return M5ConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
