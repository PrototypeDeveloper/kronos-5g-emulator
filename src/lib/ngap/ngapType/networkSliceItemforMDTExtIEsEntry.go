// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NetworkSliceItemforMDTExtIEsEntryPresent uint32

const (
	NetworkSliceItemforMDTExtIEsEntryPresentNothing NetworkSliceItemforMDTExtIEsEntryPresent = iota
	NetworkSliceItemforMDTExtIEsEntryPresentUnknown
)

type NetworkSliceItemforMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NetworkSliceItemforMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NetworkSliceItemforMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NetworkSliceItemforMDT-ExtIEsEntry", v)
}

// DecodeNetworkSliceItemforMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNetworkSliceItemforMDTExtIEsEntry(data []byte) (NetworkSliceItemforMDTExtIEsEntry, int, error) {
	var v NetworkSliceItemforMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("NetworkSliceItemforMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NetworkSliceItemforMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
