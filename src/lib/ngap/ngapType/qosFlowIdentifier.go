// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowIdentifier represents ASN.1 QosFlowIdentifier.
type QosFlowIdentifier struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:63"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowIdentifier) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowIdentifier", v)
}

// DecodeQosFlowIdentifier decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowIdentifier(data []byte) (QosFlowIdentifier, int, error) {
	var v QosFlowIdentifier
	err := ngapcodec.DecodeParameter("QosFlowIdentifier", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowIdentifier{}, 0, err
	}
	return v, len(data), err
}
