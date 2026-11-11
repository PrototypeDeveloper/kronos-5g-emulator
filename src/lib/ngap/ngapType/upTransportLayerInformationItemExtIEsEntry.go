// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UPTransportLayerInformationItemExtIEsEntryPresent uint32

const (
	UPTransportLayerInformationItemExtIEsEntryPresentNothing UPTransportLayerInformationItemExtIEsEntryPresent = iota
	UPTransportLayerInformationItemExtIEsEntryPresentCommonNetworkInstance
	UPTransportLayerInformationItemExtIEsEntryPresentUnknown
)

type UPTransportLayerInformationItemExtIEsEntry struct {
	ID                    ProtocolExtensionID
	Criticality           Criticality
	Present               UPTransportLayerInformationItemExtIEsEntryPresent
	CommonNetworkInstance CommonNetworkInstance `ngap:"id:166,criticality:ignore,presence:optional"`
	UnknownValue          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UPTransportLayerInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UPTransportLayerInformationItem-ExtIEsEntry", v)
}

// DecodeUPTransportLayerInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUPTransportLayerInformationItemExtIEsEntry(data []byte) (UPTransportLayerInformationItemExtIEsEntry, int, error) {
	var v UPTransportLayerInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("UPTransportLayerInformationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UPTransportLayerInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
