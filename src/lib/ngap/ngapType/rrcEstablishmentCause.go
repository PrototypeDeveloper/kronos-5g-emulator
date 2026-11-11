// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RRCEstablishmentCause represents ASN.1 RRCEstablishmentCause.
const (
	RRCEstablishmentCausePresentEmergency          Enumerated = 0
	RRCEstablishmentCausePresentHighPriorityAccess Enumerated = 1
	RRCEstablishmentCausePresentMtAccess           Enumerated = 2
	RRCEstablishmentCausePresentMoSignalling       Enumerated = 3
	RRCEstablishmentCausePresentMoData             Enumerated = 4
	RRCEstablishmentCausePresentMoVoiceCall        Enumerated = 5
	RRCEstablishmentCausePresentMoVideoCall        Enumerated = 6
	RRCEstablishmentCausePresentMoSMS              Enumerated = 7
	RRCEstablishmentCausePresentMpsPriorityAccess  Enumerated = 8
	RRCEstablishmentCausePresentMcsPriorityAccess  Enumerated = 9
	RRCEstablishmentCausePresentNotAvailable       Enumerated = 10
	RRCEstablishmentCausePresentMoExceptionData    Enumerated = 11
)

type RRCEstablishmentCause struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:9"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RRCEstablishmentCause) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RRCEstablishmentCause", v)
}

// DecodeRRCEstablishmentCause decodes a standalone APER value and reports the consumed octets.
func DecodeRRCEstablishmentCause(data []byte) (RRCEstablishmentCause, int, error) {
	var v RRCEstablishmentCause
	err := ngapcodec.DecodeParameter("RRCEstablishmentCause", data, &v)
	if ngapcodec.IsFatal(err) {
		return RRCEstablishmentCause{}, 0, err
	}
	return v, len(data), err
}
