// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssistanceInformationQoEMeas represents ASN.1 AssistanceInformationQoE-Meas.
type AssistanceInformationQoEMeas struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssistanceInformationQoEMeas) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssistanceInformationQoE-Meas", v)
}

// DecodeAssistanceInformationQoEMeas decodes a standalone APER value and reports the consumed octets.
func DecodeAssistanceInformationQoEMeas(data []byte) (AssistanceInformationQoEMeas, int, error) {
	var v AssistanceInformationQoEMeas
	err := ngapcodec.DecodeParameter("AssistanceInformationQoE-Meas", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssistanceInformationQoEMeas{}, 0, err
	}
	return v, len(data), err
}
