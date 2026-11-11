// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEHistoryInformationFromTheUEExtIEsEntryPresent uint32

const (
	UEHistoryInformationFromTheUEExtIEsEntryPresentNothing UEHistoryInformationFromTheUEExtIEsEntryPresent = iota
	UEHistoryInformationFromTheUEExtIEsEntryPresentUnknown
)

type UEHistoryInformationFromTheUEExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      UEHistoryInformationFromTheUEExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEHistoryInformationFromTheUEExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEHistoryInformationFromTheUE-ExtIEsEntry", v)
}

// DecodeUEHistoryInformationFromTheUEExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEHistoryInformationFromTheUEExtIEsEntry(data []byte) (UEHistoryInformationFromTheUEExtIEsEntry, int, error) {
	var v UEHistoryInformationFromTheUEExtIEsEntry
	err := ngapcodec.DecodeParameter("UEHistoryInformationFromTheUE-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEHistoryInformationFromTheUEExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
