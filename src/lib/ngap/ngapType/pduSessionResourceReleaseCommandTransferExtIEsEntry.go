// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceReleaseCommandTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceReleaseCommandTransferExtIEsEntryPresentNothing PDUSessionResourceReleaseCommandTransferExtIEsEntryPresent = iota
	PDUSessionResourceReleaseCommandTransferExtIEsEntryPresentQosFlowAdditionalInfoList
	PDUSessionResourceReleaseCommandTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceReleaseCommandTransferExtIEsEntry struct {
	ID                        ProtocolExtensionID
	Criticality               Criticality
	Present                   PDUSessionResourceReleaseCommandTransferExtIEsEntryPresent
	QosFlowAdditionalInfoList QosFlowAdditionalInfoListRelCom `ngap:"id:397,criticality:ignore,presence:optional"`
	UnknownValue              []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleaseCommandTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleaseCommandTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceReleaseCommandTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleaseCommandTransferExtIEsEntry(data []byte) (PDUSessionResourceReleaseCommandTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceReleaseCommandTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleaseCommandTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleaseCommandTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
