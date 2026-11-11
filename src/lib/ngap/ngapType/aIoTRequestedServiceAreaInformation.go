// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIoTRequestedServiceAreaInformation represents ASN.1 AIoT-RequestedServiceAreaInformation.
type AIoTRequestedServiceAreaInformation struct {
	HasRequestedReaderList   bool
	RequestedReaderList      AIoTRequestedReaderList
	HasRequestedAIoTAreaList bool
	RequestedAIoTAreaList    AIoTRequestedAIoTAreaList
	HasIEExtensions          bool
	IEExtensions             []AIoTRequestedServiceAreaInformationExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTRequestedServiceAreaInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-RequestedServiceAreaInformation", v)
}

// DecodeAIoTRequestedServiceAreaInformation decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTRequestedServiceAreaInformation(data []byte) (AIoTRequestedServiceAreaInformation, int, error) {
	var v AIoTRequestedServiceAreaInformation
	err := ngapcodec.DecodeParameter("AIoT-RequestedServiceAreaInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTRequestedServiceAreaInformation{}, 0, err
	}
	return v, len(data), err
}
