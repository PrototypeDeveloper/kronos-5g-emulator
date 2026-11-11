// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntryPresent uint32

const (
	NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntryPresentNothing NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntryPresent = iota
	NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntryPresentUnknown
)

type NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-Paging-Long-eDRX-Information-for-RRC-INACTIVE-ExtIEsEntry", v)
}

// DecodeNRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntry(data []byte) (NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntry, int, error) {
	var v NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntry
	err := ngapcodec.DecodeParameter("NR-Paging-Long-eDRX-Information-for-RRC-INACTIVE-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPagingLongEDRXInformationForRRCINACTIVEExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
