// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSDataForwardingResponseMRBItemExtIEsEntryPresent uint32

const (
	MBSDataForwardingResponseMRBItemExtIEsEntryPresentNothing MBSDataForwardingResponseMRBItemExtIEsEntryPresent = iota
	MBSDataForwardingResponseMRBItemExtIEsEntryPresentUnknown
)

type MBSDataForwardingResponseMRBItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSDataForwardingResponseMRBItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDataForwardingResponseMRBItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DataForwardingResponseMRBItem-ExtIEsEntry", v)
}

// DecodeMBSDataForwardingResponseMRBItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDataForwardingResponseMRBItemExtIEsEntry(data []byte) (MBSDataForwardingResponseMRBItemExtIEsEntry, int, error) {
	var v MBSDataForwardingResponseMRBItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-DataForwardingResponseMRBItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDataForwardingResponseMRBItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
