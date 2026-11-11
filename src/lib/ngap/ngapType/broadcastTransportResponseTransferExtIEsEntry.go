// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BroadcastTransportResponseTransferExtIEsEntryPresent uint32

const (
	BroadcastTransportResponseTransferExtIEsEntryPresentNothing BroadcastTransportResponseTransferExtIEsEntryPresent = iota
	BroadcastTransportResponseTransferExtIEsEntryPresentUnknown
)

type BroadcastTransportResponseTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      BroadcastTransportResponseTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportResponseTransfer-ExtIEsEntry", v)
}

// DecodeBroadcastTransportResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportResponseTransferExtIEsEntry(data []byte) (BroadcastTransportResponseTransferExtIEsEntry, int, error) {
	var v BroadcastTransportResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("BroadcastTransportResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
