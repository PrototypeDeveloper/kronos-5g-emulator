// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TSCTrafficCharacteristicsFeedback represents ASN.1 TSCTrafficCharacteristicsFeedback.
type TSCTrafficCharacteristicsFeedback struct {
	HasTSCFeedbackInformationDL bool
	TSCFeedbackInformationDL    TSCFeedbackInformation
	HasTSCFeedbackInformationUL bool
	TSCFeedbackInformationUL    TSCFeedbackInformation
	HasIEExtensions             bool
	IEExtensions                []TSCTrafficCharacteristicsFeedbackExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TSCTrafficCharacteristicsFeedback) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TSCTrafficCharacteristicsFeedback", v)
}

// DecodeTSCTrafficCharacteristicsFeedback decodes a standalone APER value and reports the consumed octets.
func DecodeTSCTrafficCharacteristicsFeedback(data []byte) (TSCTrafficCharacteristicsFeedback, int, error) {
	var v TSCTrafficCharacteristicsFeedback
	err := ngapcodec.DecodeParameter("TSCTrafficCharacteristicsFeedback", data, &v)
	if ngapcodec.IsFatal(err) {
		return TSCTrafficCharacteristicsFeedback{}, 0, err
	}
	return v, len(data), err
}
