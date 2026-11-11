// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSetupResponseTransfer represents ASN.1 PDUSessionResourceSetupResponseTransfer.
type PDUSessionResourceSetupResponseTransfer struct {
	DLQosFlowPerTNLInformation              QosFlowPerTNLInformation
	HasAdditionalDLQosFlowPerTNLInformation bool
	AdditionalDLQosFlowPerTNLInformation    QosFlowPerTNLInformationList
	HasSecurityResult                       bool
	SecurityResult                          SecurityResult
	HasQosFlowFailedToSetupList             bool
	QosFlowFailedToSetupList                QosFlowListWithCause
	HasIEExtensions                         bool
	IEExtensions                            []PDUSessionResourceSetupResponseTransferExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupResponseTransfer", v)
}

// DecodePDUSessionResourceSetupResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupResponseTransfer(data []byte) (PDUSessionResourceSetupResponseTransfer, int, error) {
	var v PDUSessionResourceSetupResponseTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
