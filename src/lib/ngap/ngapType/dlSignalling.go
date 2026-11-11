// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DLSignalling represents ASN.1 DL-Signalling.
const (
	DLSignallingPresentTrue Enumerated = 0
)

type DLSignalling struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLSignalling) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DL-Signalling", v) }

// DecodeDLSignalling decodes a standalone APER value and reports the consumed octets.
func DecodeDLSignalling(data []byte) (DLSignalling, int, error) {
	var v DLSignalling
	err := ngapcodec.DecodeParameter("DL-Signalling", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLSignalling{}, 0, err
	}
	return v, len(data), err
}
