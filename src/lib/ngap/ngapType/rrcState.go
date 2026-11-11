// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RRCState represents ASN.1 RRCState.
const (
	RRCStatePresentInactive  Enumerated = 0
	RRCStatePresentConnected Enumerated = 1
)

type RRCState struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RRCState) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("RRCState", v) }

// DecodeRRCState decodes a standalone APER value and reports the consumed octets.
func DecodeRRCState(data []byte) (RRCState, int, error) {
	var v RRCState
	err := ngapcodec.DecodeParameter("RRCState", data, &v)
	if ngapcodec.IsFatal(err) {
		return RRCState{}, 0, err
	}
	return v, len(data), err
}
