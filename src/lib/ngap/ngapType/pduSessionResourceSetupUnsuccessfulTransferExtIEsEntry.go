// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntryPresentNothing PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntryPresent = iota
	PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupUnsuccessfulTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSetupUnsuccessfulTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupUnsuccessfulTransferExtIEsEntry(data []byte) (PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupUnsuccessfulTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupUnsuccessfulTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
