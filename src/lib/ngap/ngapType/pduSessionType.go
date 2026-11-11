// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionType represents ASN.1 PDUSessionType.
const (
	PDUSessionTypePresentIpv4         Enumerated = 0
	PDUSessionTypePresentIpv6         Enumerated = 1
	PDUSessionTypePresentIpv4v6       Enumerated = 2
	PDUSessionTypePresentEthernet     Enumerated = 3
	PDUSessionTypePresentUnstructured Enumerated = 4
)

type PDUSessionType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:4"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionType", v)
}

// DecodePDUSessionType decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionType(data []byte) (PDUSessionType, int, error) {
	var v PDUSessionType
	err := ngapcodec.DecodeParameter("PDUSessionType", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionType{}, 0, err
	}
	return v, len(data), err
}
