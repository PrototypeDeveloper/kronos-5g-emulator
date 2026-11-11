// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BroadcastTransportFailureTransferExtIEsEntryPresent uint32

const (
	BroadcastTransportFailureTransferExtIEsEntryPresentNothing BroadcastTransportFailureTransferExtIEsEntryPresent = iota
	BroadcastTransportFailureTransferExtIEsEntryPresentUnknown
)

type BroadcastTransportFailureTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      BroadcastTransportFailureTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportFailureTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportFailureTransfer-ExtIEsEntry", v)
}

// DecodeBroadcastTransportFailureTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportFailureTransferExtIEsEntry(data []byte) (BroadcastTransportFailureTransferExtIEsEntry, int, error) {
	var v BroadcastTransportFailureTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("BroadcastTransportFailureTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportFailureTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
