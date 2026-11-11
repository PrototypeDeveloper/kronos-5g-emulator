// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyConfirmTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyConfirmTransferExtIEsEntryPresentNothing PDUSessionResourceModifyConfirmTransferExtIEsEntryPresent = iota
	PDUSessionResourceModifyConfirmTransferExtIEsEntryPresentRedundantULNGUUPTNLInformation
	PDUSessionResourceModifyConfirmTransferExtIEsEntryPresentAdditionalRedundantNGUUPTNLInformation
	PDUSessionResourceModifyConfirmTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyConfirmTransferExtIEsEntry struct {
	ID                                     ProtocolExtensionID
	Criticality                            Criticality
	Present                                PDUSessionResourceModifyConfirmTransferExtIEsEntryPresent
	RedundantULNGUUPTNLInformation         UPTransportLayerInformation         `ngap:"id:195,criticality:ignore,presence:optional"`
	AdditionalRedundantNGUUPTNLInformation UPTransportLayerInformationPairList `ngap:"id:185,criticality:ignore,presence:optional"`
	UnknownValue                           []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyConfirmTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyConfirmTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyConfirmTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyConfirmTransferExtIEsEntry(data []byte) (PDUSessionResourceModifyConfirmTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyConfirmTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyConfirmTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyConfirmTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
