// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ECNMarkingorCongestionInformationReportingRequest represents ASN.1 ECNMarkingorCongestionInformationReportingRequest.
type ECNMarkingorCongestionInformationReportingRequestPresent uint32

const (
	ECNMarkingorCongestionInformationReportingRequestPresentNothing ECNMarkingorCongestionInformationReportingRequestPresent = iota
	ECNMarkingorCongestionInformationReportingRequestPresentECNMarkingAtRANRequest
	ECNMarkingorCongestionInformationReportingRequestPresentECNMarkingAtUPFRequest
	ECNMarkingorCongestionInformationReportingRequestPresentCongestionInformationRequest
	ECNMarkingorCongestionInformationReportingRequestPresentChoiceExtensions
)

type ECNMarkingorCongestionInformationReportingRequest struct {
	Present                      ECNMarkingorCongestionInformationReportingRequestPresent
	ECNMarkingAtRANRequest       ECNMarkingAtRANRequest
	ECNMarkingAtUPFRequest       ECNMarkingAtUPFRequest
	CongestionInformationRequest CongestionInformationRequest
	ChoiceExtensions             ECNMarkingorCongestionInformationReportingRequestExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ECNMarkingorCongestionInformationReportingRequest) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ECNMarkingorCongestionInformationReportingRequest", v)
}

// DecodeECNMarkingorCongestionInformationReportingRequest decodes a standalone APER value and reports the consumed octets.
func DecodeECNMarkingorCongestionInformationReportingRequest(data []byte) (ECNMarkingorCongestionInformationReportingRequest, int, error) {
	var v ECNMarkingorCongestionInformationReportingRequest
	err := ngapcodec.DecodeParameter("ECNMarkingorCongestionInformationReportingRequest", data, &v)
	if ngapcodec.IsFatal(err) {
		return ECNMarkingorCongestionInformationReportingRequest{}, 0, err
	}
	return v, len(data), err
}
