// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UPTransportLayerInformationPairItemExtIEsEntryPresent uint32

const (
	UPTransportLayerInformationPairItemExtIEsEntryPresentNothing UPTransportLayerInformationPairItemExtIEsEntryPresent = iota
	UPTransportLayerInformationPairItemExtIEsEntryPresentUnknown
)

type UPTransportLayerInformationPairItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UPTransportLayerInformationPairItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UPTransportLayerInformationPairItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UPTransportLayerInformationPairItem-ExtIEsEntry", v)
}

// DecodeUPTransportLayerInformationPairItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUPTransportLayerInformationPairItemExtIEsEntry(data []byte) (UPTransportLayerInformationPairItemExtIEsEntry, int, error) {
	var v UPTransportLayerInformationPairItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UPTransportLayerInformationPairItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UPTransportLayerInformationPairItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
