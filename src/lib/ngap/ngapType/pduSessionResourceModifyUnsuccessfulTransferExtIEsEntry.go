// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntryPresentNothing PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntryPresent = iota
	PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyUnsuccessfulTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyUnsuccessfulTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyUnsuccessfulTransferExtIEsEntry(data []byte) (PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyUnsuccessfulTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyUnsuccessfulTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
