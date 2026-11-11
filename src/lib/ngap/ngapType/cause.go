// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// Cause represents ASN.1 Cause.
type CausePresent uint32

const (
	CausePresentNothing CausePresent = iota
	CausePresentRadioNetwork
	CausePresentTransport
	CausePresentNas
	CausePresentProtocol
	CausePresentMisc
	CausePresentChoiceExtensions
)

type Cause struct {
	Present          CausePresent
	RadioNetwork     CauseRadioNetwork
	Transport        CauseTransport
	Nas              CauseNas
	Protocol         CauseProtocol
	Misc             CauseMisc
	ChoiceExtensions CauseExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Cause) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("Cause", v) }

// DecodeCause decodes a standalone APER value and reports the consumed octets.
func DecodeCause(data []byte) (Cause, int, error) {
	var v Cause
	err := ngapcodec.DecodeParameter("Cause", data, &v)
	if ngapcodec.IsFatal(err) {
		return Cause{}, 0, err
	}
	return v, len(data), err
}
