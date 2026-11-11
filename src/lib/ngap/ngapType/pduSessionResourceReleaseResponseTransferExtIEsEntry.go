// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceReleaseResponseTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceReleaseResponseTransferExtIEsEntryPresentNothing PDUSessionResourceReleaseResponseTransferExtIEsEntryPresent = iota
	PDUSessionResourceReleaseResponseTransferExtIEsEntryPresentSecondaryRATUsageInformation
	PDUSessionResourceReleaseResponseTransferExtIEsEntryPresentQosFlowAdditionalInfoList
	PDUSessionResourceReleaseResponseTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceReleaseResponseTransferExtIEsEntry struct {
	ID                           ProtocolExtensionID
	Criticality                  Criticality
	Present                      PDUSessionResourceReleaseResponseTransferExtIEsEntryPresent
	SecondaryRATUsageInformation SecondaryRATUsageInformation    `ngap:"id:144,criticality:ignore,presence:optional"`
	QosFlowAdditionalInfoList    QosFlowAdditionalInfoListRelRes `ngap:"id:397,criticality:ignore,presence:optional"`
	UnknownValue                 []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleaseResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleaseResponseTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceReleaseResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleaseResponseTransferExtIEsEntry(data []byte) (PDUSessionResourceReleaseResponseTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceReleaseResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleaseResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleaseResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
