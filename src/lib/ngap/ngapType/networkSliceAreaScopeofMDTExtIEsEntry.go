// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NetworkSliceAreaScopeofMDTExtIEsEntryPresent uint32

const (
	NetworkSliceAreaScopeofMDTExtIEsEntryPresentNothing NetworkSliceAreaScopeofMDTExtIEsEntryPresent = iota
	NetworkSliceAreaScopeofMDTExtIEsEntryPresentUnknown
)

type NetworkSliceAreaScopeofMDTExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NetworkSliceAreaScopeofMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NetworkSliceAreaScopeofMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NetworkSliceAreaScopeofMDT-ExtIEsEntry", v)
}

// DecodeNetworkSliceAreaScopeofMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNetworkSliceAreaScopeofMDTExtIEsEntry(data []byte) (NetworkSliceAreaScopeofMDTExtIEsEntry, int, error) {
	var v NetworkSliceAreaScopeofMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("NetworkSliceAreaScopeofMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NetworkSliceAreaScopeofMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
