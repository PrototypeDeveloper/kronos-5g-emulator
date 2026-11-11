// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NBIoTPagingEDRXInfoExtIEsEntryPresent uint32

const (
	NBIoTPagingEDRXInfoExtIEsEntryPresentNothing NBIoTPagingEDRXInfoExtIEsEntryPresent = iota
	NBIoTPagingEDRXInfoExtIEsEntryPresentUnknown
)

type NBIoTPagingEDRXInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NBIoTPagingEDRXInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NBIoTPagingEDRXInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NB-IoT-Paging-eDRXInfo-ExtIEsEntry", v)
}

// DecodeNBIoTPagingEDRXInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNBIoTPagingEDRXInfoExtIEsEntry(data []byte) (NBIoTPagingEDRXInfoExtIEsEntry, int, error) {
	var v NBIoTPagingEDRXInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("NB-IoT-Paging-eDRXInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NBIoTPagingEDRXInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
