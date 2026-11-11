// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceNotifyTransfer represents ASN.1 PDUSessionResourceNotifyTransfer.
type PDUSessionResourceNotifyTransfer struct {
	HasQosFlowNotifyList   bool
	QosFlowNotifyList      QosFlowNotifyList
	HasQosFlowReleasedList bool
	QosFlowReleasedList    QosFlowListWithCause
	HasIEExtensions        bool
	IEExtensions           []PDUSessionResourceNotifyTransferExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceNotifyTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceNotifyTransfer", v)
}

// DecodePDUSessionResourceNotifyTransfer decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceNotifyTransfer(data []byte) (PDUSessionResourceNotifyTransfer, int, error) {
	var v PDUSessionResourceNotifyTransfer
	err := ngapcodec.DecodeParameter("PDUSessionResourceNotifyTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceNotifyTransfer{}, 0, err
	}
	return v, len(data), err
}
