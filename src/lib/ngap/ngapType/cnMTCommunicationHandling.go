// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CNMTCommunicationHandling represents ASN.1 CN-MT-CommunicationHandling.
const (
	CNMTCommunicationHandlingPresentSupported Enumerated = 0
)

type CNMTCommunicationHandling struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CNMTCommunicationHandling) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CN-MT-CommunicationHandling", v)
}

// DecodeCNMTCommunicationHandling decodes a standalone APER value and reports the consumed octets.
func DecodeCNMTCommunicationHandling(data []byte) (CNMTCommunicationHandling, int, error) {
	var v CNMTCommunicationHandling
	err := ngapcodec.DecodeParameter("CN-MT-CommunicationHandling", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNMTCommunicationHandling{}, 0, err
	}
	return v, len(data), err
}
