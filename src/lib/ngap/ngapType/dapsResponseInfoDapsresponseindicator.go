// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DAPSResponseInfoDapsresponseindicator represents ASN.1 DAPSResponseInfoDapsresponseindicator.
const (
	DAPSResponseInfoDapsresponseindicatorPresentDapsHoAccepted    Enumerated = 0
	DAPSResponseInfoDapsresponseindicatorPresentDapsHoNotAccepted Enumerated = 1
)

type DAPSResponseInfoDapsresponseindicator struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v DAPSResponseInfoDapsresponseindicator) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:DAPSResponseInfoDapsresponseindicator", v)
}

// DecodeDAPSResponseInfoDapsresponseindicator decodes a standalone APER value and reports consumed octets.
func DecodeDAPSResponseInfoDapsresponseindicator(data []byte) (DAPSResponseInfoDapsresponseindicator, int, error) {
	var v DAPSResponseInfoDapsresponseindicator
	err := ngapcodec.DecodeParameter("inline:DAPSResponseInfoDapsresponseindicator", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSResponseInfoDapsresponseindicator{}, 0, err
	}
	return v, len(data), err
}
