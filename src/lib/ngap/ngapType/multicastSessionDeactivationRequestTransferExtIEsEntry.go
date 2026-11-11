// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MulticastSessionDeactivationRequestTransferExtIEsEntryPresent uint32

const (
	MulticastSessionDeactivationRequestTransferExtIEsEntryPresentNothing MulticastSessionDeactivationRequestTransferExtIEsEntryPresent = iota
	MulticastSessionDeactivationRequestTransferExtIEsEntryPresentUnknown
)

type MulticastSessionDeactivationRequestTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MulticastSessionDeactivationRequestTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionDeactivationRequestTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionDeactivationRequestTransfer-ExtIEsEntry", v)
}

// DecodeMulticastSessionDeactivationRequestTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionDeactivationRequestTransferExtIEsEntry(data []byte) (MulticastSessionDeactivationRequestTransferExtIEsEntry, int, error) {
	var v MulticastSessionDeactivationRequestTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MulticastSessionDeactivationRequestTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionDeactivationRequestTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
