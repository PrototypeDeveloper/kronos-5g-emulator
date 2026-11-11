// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyConfirmTransfer represents ASN.1 PDUSessionResourceModifyConfirmTransfer.
type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList         QosFlowModifyConfirmList
	ULNGUUPTNLInformation            UPTransportLayerInformation
	HasAdditionalNGUUPTNLInformation bool
	AdditionalNGUUPTNLInformation    UPTransportLayerInformationPairList
	HasQosFlowFailedToModifyList     bool
	QosFlowFailedToModifyList        QosFlowListWithCause
	HasIEExtensions                  bool
	IEExtensions                     []PDUSessionResourceModifyConfirmTransferExtIEsEntry
	UnknownExtensions                []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyConfirmTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyConfirmTransfer", v)
}

// DecodePDUSessionResourceModifyConfirmTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyConfirmTransfer(data []byte) (PDUSessionResourceModifyConfirmTransfer, int, error) {
	var v PDUSessionResourceModifyConfirmTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyConfirmTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyConfirmTransfer{}, 0, err
	}
	return v, len(data), err
}
