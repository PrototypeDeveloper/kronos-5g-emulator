// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUsetQoSInformationPduSetIntegratedHandlingInformation represents ASN.1 PDUsetQoSInformationPduSetIntegratedHandlingInformation.
const (
	PDUsetQoSInformationPduSetIntegratedHandlingInformationPresentTrue  Enumerated = 0
	PDUsetQoSInformationPduSetIntegratedHandlingInformationPresentFalse Enumerated = 1
)

type PDUsetQoSInformationPduSetIntegratedHandlingInformation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v PDUsetQoSInformationPduSetIntegratedHandlingInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:PDUsetQoSInformationPduSetIntegratedHandlingInformation", v)
}

// DecodePDUsetQoSInformationPduSetIntegratedHandlingInformation decodes a standalone APER value and reports consumed octets.
func DecodePDUsetQoSInformationPduSetIntegratedHandlingInformation(data []byte) (PDUsetQoSInformationPduSetIntegratedHandlingInformation, int, error) {
	var v PDUsetQoSInformationPduSetIntegratedHandlingInformation
	err := ngapcodec.DecodeParameter("inline:PDUsetQoSInformationPduSetIntegratedHandlingInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUsetQoSInformationPduSetIntegratedHandlingInformation{}, 0, err
	}
	return v, len(data), err
}
