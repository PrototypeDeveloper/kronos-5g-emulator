// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MulticastSessionActivationRequestTransferExtIEsEntryPresent uint32

const (
	MulticastSessionActivationRequestTransferExtIEsEntryPresentNothing MulticastSessionActivationRequestTransferExtIEsEntryPresent = iota
	MulticastSessionActivationRequestTransferExtIEsEntryPresentUnknown
)

type MulticastSessionActivationRequestTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MulticastSessionActivationRequestTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MulticastSessionActivationRequestTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MulticastSessionActivationRequestTransfer-ExtIEsEntry", v)
}

// DecodeMulticastSessionActivationRequestTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMulticastSessionActivationRequestTransferExtIEsEntry(data []byte) (MulticastSessionActivationRequestTransferExtIEsEntry, int, error) {
	var v MulticastSessionActivationRequestTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("MulticastSessionActivationRequestTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MulticastSessionActivationRequestTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
