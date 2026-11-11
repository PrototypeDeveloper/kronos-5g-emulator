// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTReaderReportItemExtIEsEntryPresent uint32

const (
	AIoTReaderReportItemExtIEsEntryPresentNothing AIoTReaderReportItemExtIEsEntryPresent = iota
	AIoTReaderReportItemExtIEsEntryPresentUnknown
)

type AIoTReaderReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AIoTReaderReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTReaderReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-ReaderReportItem-ExtIEsEntry", v)
}

// DecodeAIoTReaderReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTReaderReportItemExtIEsEntry(data []byte) (AIoTReaderReportItemExtIEsEntry, int, error) {
	var v AIoTReaderReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoT-ReaderReportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTReaderReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
