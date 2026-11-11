// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEContextSuspendRequestTransferExtIEsEntryPresent uint32

const (
	UEContextSuspendRequestTransferExtIEsEntryPresentNothing UEContextSuspendRequestTransferExtIEsEntryPresent = iota
	UEContextSuspendRequestTransferExtIEsEntryPresentUnknown
)

type UEContextSuspendRequestTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEContextSuspendRequestTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEContextSuspendRequestTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEContextSuspendRequestTransfer-ExtIEsEntry", v)
}

// DecodeUEContextSuspendRequestTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEContextSuspendRequestTransferExtIEsEntry(data []byte) (UEContextSuspendRequestTransferExtIEsEntry, int, error) {
	var v UEContextSuspendRequestTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("UEContextSuspendRequestTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEContextSuspendRequestTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
