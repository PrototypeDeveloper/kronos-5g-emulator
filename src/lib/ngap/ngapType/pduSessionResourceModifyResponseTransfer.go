// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyResponseTransfer represents ASN.1 PDUSessionResourceModifyResponseTransfer.
type PDUSessionResourceModifyResponseTransfer struct {
	HasDLNGUUPTNLInformation                bool
	DLNGUUPTNLInformation                   UPTransportLayerInformation
	HasULNGUUPTNLInformation                bool
	ULNGUUPTNLInformation                   UPTransportLayerInformation
	HasQosFlowAddOrModifyResponseList       bool
	QosFlowAddOrModifyResponseList          QosFlowAddOrModifyResponseList
	HasAdditionalDLQosFlowPerTNLInformation bool
	AdditionalDLQosFlowPerTNLInformation    QosFlowPerTNLInformationList
	HasQosFlowFailedToAddOrModifyList       bool
	QosFlowFailedToAddOrModifyList          QosFlowListWithCause
	HasIEExtensions                         bool
	IEExtensions                            []PDUSessionResourceModifyResponseTransferExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyResponseTransfer", v)
}

// DecodePDUSessionResourceModifyResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyResponseTransfer(data []byte) (PDUSessionResourceModifyResponseTransfer, int, error) {
	var v PDUSessionResourceModifyResponseTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
