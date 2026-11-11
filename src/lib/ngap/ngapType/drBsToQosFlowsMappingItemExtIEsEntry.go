// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBsToQosFlowsMappingItemExtIEsEntryPresent uint32

const (
	DRBsToQosFlowsMappingItemExtIEsEntryPresentNothing DRBsToQosFlowsMappingItemExtIEsEntryPresent = iota
	DRBsToQosFlowsMappingItemExtIEsEntryPresentDAPSRequestInfo
	DRBsToQosFlowsMappingItemExtIEsEntryPresentUnknown
)

type DRBsToQosFlowsMappingItemExtIEsEntry struct {
	ID              ProtocolExtensionID
	Criticality     Criticality
	Present         DRBsToQosFlowsMappingItemExtIEsEntryPresent
	DAPSRequestInfo DAPSRequestInfo `ngap:"id:266,criticality:ignore,presence:optional"`
	UnknownValue    []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsToQosFlowsMappingItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsToQosFlowsMappingItem-ExtIEsEntry", v)
}

// DecodeDRBsToQosFlowsMappingItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsToQosFlowsMappingItemExtIEsEntry(data []byte) (DRBsToQosFlowsMappingItemExtIEsEntry, int, error) {
	var v DRBsToQosFlowsMappingItemExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBsToQosFlowsMappingItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsToQosFlowsMappingItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
