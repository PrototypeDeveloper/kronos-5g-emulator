// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CoreNetworkAssistanceInformationForInactive represents ASN.1 CoreNetworkAssistanceInformationForInactive.
type CoreNetworkAssistanceInformationForInactive struct {
	UEIdentityIndexValue            UEIdentityIndexValue
	HasUESpecificDRX                bool
	UESpecificDRX                   PagingDRX
	PeriodicRegistrationUpdateTimer PeriodicRegistrationUpdateTimer
	HasMICOModeIndication           bool
	MICOModeIndication              MICOModeIndication
	TAIListForInactive              TAIListForInactive
	HasExpectedUEBehaviour          bool
	ExpectedUEBehaviour             ExpectedUEBehaviour
	HasIEExtensions                 bool
	IEExtensions                    []CoreNetworkAssistanceInformationForInactiveExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CoreNetworkAssistanceInformationForInactive) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CoreNetworkAssistanceInformationForInactive", v)
}

// DecodeCoreNetworkAssistanceInformationForInactive decodes a standalone APER value and reports the consumed octets.
func DecodeCoreNetworkAssistanceInformationForInactive(data []byte) (CoreNetworkAssistanceInformationForInactive, int, error) {
	var v CoreNetworkAssistanceInformationForInactive
	err := ngapcodec.DecodeParameter("CoreNetworkAssistanceInformationForInactive", data, &v)
	if ngapcodec.IsFatal(err) {
		return CoreNetworkAssistanceInformationForInactive{}, 0, err
	}
	return v, len(data), err
}
