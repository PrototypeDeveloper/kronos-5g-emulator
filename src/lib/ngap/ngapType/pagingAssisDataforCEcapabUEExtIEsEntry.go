// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PagingAssisDataforCEcapabUEExtIEsEntryPresent uint32

const (
	PagingAssisDataforCEcapabUEExtIEsEntryPresentNothing PagingAssisDataforCEcapabUEExtIEsEntryPresent = iota
	PagingAssisDataforCEcapabUEExtIEsEntryPresentUnknown
)

type PagingAssisDataforCEcapabUEExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PagingAssisDataforCEcapabUEExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingAssisDataforCEcapabUEExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingAssisDataforCEcapabUE-ExtIEsEntry", v)
}

// DecodePagingAssisDataforCEcapabUEExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePagingAssisDataforCEcapabUEExtIEsEntry(data []byte) (PagingAssisDataforCEcapabUEExtIEsEntry, int, error) {
	var v PagingAssisDataforCEcapabUEExtIEsEntry
	err := ngapcodec.DecodeParameter("PagingAssisDataforCEcapabUE-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingAssisDataforCEcapabUEExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
