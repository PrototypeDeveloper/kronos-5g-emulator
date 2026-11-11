// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceNotifyTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceNotifyTransferExtIEsEntryPresentNothing PDUSessionResourceNotifyTransferExtIEsEntryPresent = iota
	PDUSessionResourceNotifyTransferExtIEsEntryPresentSecondaryRATUsageInformation
	PDUSessionResourceNotifyTransferExtIEsEntryPresentQosFlowFeedbackList
	PDUSessionResourceNotifyTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceNotifyTransferExtIEsEntry struct {
	ID                           ProtocolExtensionID
	Criticality                  Criticality
	Present                      PDUSessionResourceNotifyTransferExtIEsEntryPresent
	SecondaryRATUsageInformation SecondaryRATUsageInformation `ngap:"id:144,criticality:ignore,presence:optional"`
	QosFlowFeedbackList          QosFlowFeedbackList          `ngap:"id:278,criticality:ignore,presence:optional"`
	UnknownValue                 []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceNotifyTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceNotifyTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceNotifyTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceNotifyTransferExtIEsEntry(data []byte) (PDUSessionResourceNotifyTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceNotifyTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceNotifyTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceNotifyTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
