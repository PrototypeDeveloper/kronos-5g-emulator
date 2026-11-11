// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSMappingandDataForwardingRequestItemExtIEsEntryPresent uint32

const (
	MBSMappingandDataForwardingRequestItemExtIEsEntryPresentNothing MBSMappingandDataForwardingRequestItemExtIEsEntryPresent = iota
	MBSMappingandDataForwardingRequestItemExtIEsEntryPresentUnknown
)

type MBSMappingandDataForwardingRequestItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSMappingandDataForwardingRequestItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSMappingandDataForwardingRequestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-MappingandDataForwardingRequestItem-ExtIEsEntry", v)
}

// DecodeMBSMappingandDataForwardingRequestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSMappingandDataForwardingRequestItemExtIEsEntry(data []byte) (MBSMappingandDataForwardingRequestItemExtIEsEntry, int, error) {
	var v MBSMappingandDataForwardingRequestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-MappingandDataForwardingRequestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSMappingandDataForwardingRequestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
