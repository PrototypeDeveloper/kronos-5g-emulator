// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DataForwardingResponseDRBItemExtIEsEntryPresent uint32

const (
	DataForwardingResponseDRBItemExtIEsEntryPresentNothing DataForwardingResponseDRBItemExtIEsEntryPresent = iota
	DataForwardingResponseDRBItemExtIEsEntryPresentUnknown
)

type DataForwardingResponseDRBItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DataForwardingResponseDRBItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataForwardingResponseDRBItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataForwardingResponseDRBItem-ExtIEsEntry", v)
}

// DecodeDataForwardingResponseDRBItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDataForwardingResponseDRBItemExtIEsEntry(data []byte) (DataForwardingResponseDRBItemExtIEsEntry, int, error) {
	var v DataForwardingResponseDRBItemExtIEsEntry
	err := ngapcodec.DecodeParameter("DataForwardingResponseDRBItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataForwardingResponseDRBItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
