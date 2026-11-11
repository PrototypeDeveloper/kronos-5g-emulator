// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type M4ConfigurationExtIEsEntryPresent uint32

const (
	M4ConfigurationExtIEsEntryPresentNothing M4ConfigurationExtIEsEntryPresent = iota
	M4ConfigurationExtIEsEntryPresentM4ReportAmount
	M4ConfigurationExtIEsEntryPresentUnknown
)

type M4ConfigurationExtIEsEntry struct {
	ID             ProtocolExtensionID
	Criticality    Criticality
	Present        M4ConfigurationExtIEsEntryPresent
	M4ReportAmount M4ReportAmountMDT `ngap:"id:336,criticality:ignore,presence:optional"`
	UnknownValue   []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M4ConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M4Configuration-ExtIEsEntry", v)
}

// DecodeM4ConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeM4ConfigurationExtIEsEntry(data []byte) (M4ConfigurationExtIEsEntry, int, error) {
	var v M4ConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("M4Configuration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return M4ConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
