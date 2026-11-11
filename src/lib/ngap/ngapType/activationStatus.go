// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ActivationStatus represents ASN.1 ActivationStatus.
const (
	ActivationStatusPresentActive    Enumerated = 0
	ActivationStatusPresentNotActive Enumerated = 1
)

type ActivationStatus struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ActivationStatus) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ActivationStatus", v)
}

// DecodeActivationStatus decodes a standalone APER value and reports the consumed octets.
func DecodeActivationStatus(data []byte) (ActivationStatus, int, error) {
	var v ActivationStatus
	err := ngapcodec.DecodeParameter("ActivationStatus", data, &v)
	if ngapcodec.IsFatal(err) {
		return ActivationStatus{}, 0, err
	}
	return v, len(data), err
}
