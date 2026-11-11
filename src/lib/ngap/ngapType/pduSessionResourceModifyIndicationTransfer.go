// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyIndicationTransfer represents ASN.1 PDUSessionResourceModifyIndicationTransfer.
type PDUSessionResourceModifyIndicationTransfer struct {
	DLQosFlowPerTNLInformation              QosFlowPerTNLInformation
	HasAdditionalDLQosFlowPerTNLInformation bool
	AdditionalDLQosFlowPerTNLInformation    QosFlowPerTNLInformationList
	HasIEExtensions                         bool
	IEExtensions                            []PDUSessionResourceModifyIndicationTransferExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyIndicationTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyIndicationTransfer", v)
}

// DecodePDUSessionResourceModifyIndicationTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyIndicationTransfer(data []byte) (PDUSessionResourceModifyIndicationTransfer, int, error) {
	var v PDUSessionResourceModifyIndicationTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyIndicationTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyIndicationTransfer{}, 0, err
	}
	return v, len(data), err
}
