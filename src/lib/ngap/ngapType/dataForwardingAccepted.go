// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DataForwardingAccepted represents ASN.1 DataForwardingAccepted.
const (
	DataForwardingAcceptedPresentDataForwardingAccepted Enumerated = 0
)

type DataForwardingAccepted struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataForwardingAccepted) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataForwardingAccepted", v)
}

// DecodeDataForwardingAccepted decodes a standalone APER value and reports the consumed octets.
func DecodeDataForwardingAccepted(data []byte) (DataForwardingAccepted, int, error) {
	var v DataForwardingAccepted
	err := ngapcodec.DecodeParameter("DataForwardingAccepted", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataForwardingAccepted{}, 0, err
	}
	return v, len(data), err
}
