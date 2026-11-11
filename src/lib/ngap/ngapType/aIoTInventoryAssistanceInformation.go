// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTInventoryAssistanceInformation represents ASN.1 AIoT-InventoryAssistanceInformation.
type AIoTInventoryAssistanceInformation struct {
	ExpectedD2RMsgSize              int64 `aper:"valueExt,valueLB:1,valueUB:256"`
	HasApproximateNoofTargetDevices bool
	ApproximateNoofTargetDevices    int64 `aper:"valueExt,valueLB:1,valueUB:65535"`
	HasTimeInterval                 bool
	TimeInterval                    int64 `aper:"valueExt,valueLB:1,valueUB:256"`
	HasIEExtensions                 bool
	IEExtensions                    []AIoTInventoryAssistanceInformationExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTInventoryAssistanceInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-InventoryAssistanceInformation", v)
}

// DecodeAIoTInventoryAssistanceInformation decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTInventoryAssistanceInformation(data []byte) (AIoTInventoryAssistanceInformation, int, error) {
	var v AIoTInventoryAssistanceInformation
	err := ngapcodec.DecodeParameter("AIoT-InventoryAssistanceInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTInventoryAssistanceInformation{}, 0, err
	}
	return v, len(data), err
}
