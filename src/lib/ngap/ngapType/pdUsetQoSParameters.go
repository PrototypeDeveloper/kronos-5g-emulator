// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUsetQoSParameters represents ASN.1 PDUsetQoSParameters.
type PDUsetQoSParameters struct {
	HasUlPDUSetQoSInformation bool
	UlPDUSetQoSInformation    PDUsetQoSInformation
	HasDlPDUSetQoSInformation bool
	DlPDUSetQoSInformation    PDUsetQoSInformation
	HasIEExtensions           bool
	IEExtensions              []PDUsetQoSParametersExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUsetQoSParameters) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUsetQoSParameters", v)
}

// DecodePDUsetQoSParameters decodes a standalone APER value and reports the consumed octets.
func DecodePDUsetQoSParameters(data []byte) (PDUsetQoSParameters, int, error) {
	var v PDUsetQoSParameters
	err := ngapcodec.DecodeParameter("PDUsetQoSParameters", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUsetQoSParameters{}, 0, err
	}
	return v, len(data), err
}
