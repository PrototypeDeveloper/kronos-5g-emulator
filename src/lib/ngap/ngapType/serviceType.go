// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ServiceType represents ASN.1 ServiceType.
const (
	ServiceTypePresentStreaming Enumerated = 0
	ServiceTypePresentMTSI      Enumerated = 1
	ServiceTypePresentVR        Enumerated = 2
)

type ServiceType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ServiceType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ServiceType", v) }

// DecodeServiceType decodes a standalone APER value and reports the consumed octets.
func DecodeServiceType(data []byte) (ServiceType, int, error) {
	var v ServiceType
	err := ngapcodec.DecodeParameter("ServiceType", data, &v)
	if ngapcodec.IsFatal(err) {
		return ServiceType{}, 0, err
	}
	return v, len(data), err
}
