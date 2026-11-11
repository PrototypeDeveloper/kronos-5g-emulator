// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BroadcastTransportRequestTransferExtIEsEntryPresent uint32

const (
	BroadcastTransportRequestTransferExtIEsEntryPresentNothing BroadcastTransportRequestTransferExtIEsEntryPresent = iota
	BroadcastTransportRequestTransferExtIEsEntryPresentUnknown
)

type BroadcastTransportRequestTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      BroadcastTransportRequestTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastTransportRequestTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastTransportRequestTransfer-ExtIEsEntry", v)
}

// DecodeBroadcastTransportRequestTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastTransportRequestTransferExtIEsEntry(data []byte) (BroadcastTransportRequestTransferExtIEsEntry, int, error) {
	var v BroadcastTransportRequestTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("BroadcastTransportRequestTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastTransportRequestTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
