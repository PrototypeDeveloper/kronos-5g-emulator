// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CauseProtocol represents ASN.1 CauseProtocol.
const (
	CauseProtocolPresentTransferSyntaxError                          Enumerated = 0
	CauseProtocolPresentAbstractSyntaxErrorReject                    Enumerated = 1
	CauseProtocolPresentAbstractSyntaxErrorIgnoreAndNotify           Enumerated = 2
	CauseProtocolPresentMessageNotCompatibleWithReceiverState        Enumerated = 3
	CauseProtocolPresentSemanticError                                Enumerated = 4
	CauseProtocolPresentAbstractSyntaxErrorFalselyConstructedMessage Enumerated = 5
	CauseProtocolPresentUnspecified                                  Enumerated = 6
)

type CauseProtocol struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:6"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CauseProtocol) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CauseProtocol", v) }

// DecodeCauseProtocol decodes a standalone APER value and reports the consumed octets.
func DecodeCauseProtocol(data []byte) (CauseProtocol, int, error) {
	var v CauseProtocol
	err := ngapcodec.DecodeParameter("CauseProtocol", data, &v)
	if ngapcodec.IsFatal(err) {
		return CauseProtocol{}, 0, err
	}
	return v, len(data), err
}
