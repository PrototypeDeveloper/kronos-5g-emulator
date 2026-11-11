// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type HandoverPreparationUnsuccessfulTransferExtIEsEntryPresent uint32

const (
	HandoverPreparationUnsuccessfulTransferExtIEsEntryPresentNothing HandoverPreparationUnsuccessfulTransferExtIEsEntryPresent = iota
	HandoverPreparationUnsuccessfulTransferExtIEsEntryPresentUnknown
)

type HandoverPreparationUnsuccessfulTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      HandoverPreparationUnsuccessfulTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverPreparationUnsuccessfulTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HandoverPreparationUnsuccessfulTransfer-ExtIEsEntry", v)
}

// DecodeHandoverPreparationUnsuccessfulTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverPreparationUnsuccessfulTransferExtIEsEntry(data []byte) (HandoverPreparationUnsuccessfulTransferExtIEsEntry, int, error) {
	var v HandoverPreparationUnsuccessfulTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("HandoverPreparationUnsuccessfulTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverPreparationUnsuccessfulTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
