// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UPTransportLayerInformationExtIEsEntryPresent uint32

const (
	UPTransportLayerInformationExtIEsEntryPresentNothing UPTransportLayerInformationExtIEsEntryPresent = iota
	UPTransportLayerInformationExtIEsEntryPresentUnknown
)

type UPTransportLayerInformationExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      UPTransportLayerInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UPTransportLayerInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UPTransportLayerInformation-ExtIEsEntry", v)
}

// DecodeUPTransportLayerInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUPTransportLayerInformationExtIEsEntry(data []byte) (UPTransportLayerInformationExtIEsEntry, int, error) {
	var v UPTransportLayerInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("UPTransportLayerInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UPTransportLayerInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
