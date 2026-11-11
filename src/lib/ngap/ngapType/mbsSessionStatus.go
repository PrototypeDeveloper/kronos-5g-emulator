// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionStatus represents ASN.1 MBSSessionStatus.
const (
	MBSSessionStatusPresentActivated   Enumerated = 0
	MBSSessionStatusPresentDeactivated Enumerated = 1
)

type MBSSessionStatus struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionStatus) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionStatus", v)
}

// DecodeMBSSessionStatus decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionStatus(data []byte) (MBSSessionStatus, int, error) {
	var v MBSSessionStatus
	err := ngapcodec.DecodeParameter("MBSSessionStatus", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionStatus{}, 0, err
	}
	return v, len(data), err
}
