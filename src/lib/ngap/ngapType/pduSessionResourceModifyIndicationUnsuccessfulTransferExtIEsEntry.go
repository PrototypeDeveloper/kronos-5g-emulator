// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntryPresentNothing PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntryPresent = iota
	PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyIndicationUnsuccessfulTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntry(data []byte) (PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyIndicationUnsuccessfulTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
