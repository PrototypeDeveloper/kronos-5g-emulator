// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PrivateIEID represents ASN.1 PrivateIE-ID.
type PrivateIEIDPresent uint32

const (
	PrivateIEIDPresentNothing PrivateIEIDPresent = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

type PrivateIEID struct {
	Present PrivateIEIDPresent
	Local   int64 `aper:"valueLB:0,valueUB:65535"`
	Global  []uint64
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PrivateIEID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("PrivateIE-ID", v) }

// DecodePrivateIEID decodes a standalone APER value and reports the consumed octets.
func DecodePrivateIEID(data []byte) (PrivateIEID, int, error) {
	var v PrivateIEID
	err := ngapcodec.DecodeParameter("PrivateIE-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return PrivateIEID{}, 0, err
	}
	return v, len(data), err
}
