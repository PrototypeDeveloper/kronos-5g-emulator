// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PagingCauseIndicationForVoiceService represents ASN.1 PagingCauseIndicationForVoiceService.
const (
	PagingCauseIndicationForVoiceServicePresentSupported Enumerated = 0
)

type PagingCauseIndicationForVoiceService struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PagingCauseIndicationForVoiceService) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PagingCauseIndicationForVoiceService", v)
}

// DecodePagingCauseIndicationForVoiceService decodes a standalone APER value and reports the consumed octets.
func DecodePagingCauseIndicationForVoiceService(data []byte) (PagingCauseIndicationForVoiceService, int, error) {
	var v PagingCauseIndicationForVoiceService
	err := ngapcodec.DecodeParameter("PagingCauseIndicationForVoiceService", data, &v)
	if ngapcodec.IsFatal(err) {
		return PagingCauseIndicationForVoiceService{}, 0, err
	}
	return v, len(data), err
}
