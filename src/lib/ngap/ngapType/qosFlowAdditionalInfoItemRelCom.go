// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAdditionalInfoItemRelCom represents ASN.1 QosFlowAdditionalInfoItemRelCom.
type QosFlowAdditionalInfoItemRelCom struct {
	QosFlowIdentifier      QosFlowIdentifier
	HasDownlinkTLContainer bool
	DownlinkTLContainer    TLContainer
	HasIEExtensions        bool
	IEExtensions           []QosFlowAdditionalInfoItemRelComExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAdditionalInfoItemRelCom) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAdditionalInfoItemRelCom", v)
}

// DecodeQosFlowAdditionalInfoItemRelCom decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAdditionalInfoItemRelCom(data []byte) (QosFlowAdditionalInfoItemRelCom, int, error) {
	var v QosFlowAdditionalInfoItemRelCom
	err := ngapcodec.DecodeParameter("QosFlowAdditionalInfoItemRelCom", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAdditionalInfoItemRelCom{}, 0, err
	}
	return v, len(data), err
}
