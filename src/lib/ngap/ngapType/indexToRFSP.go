// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IndexToRFSP represents ASN.1 IndexToRFSP.
type IndexToRFSP struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IndexToRFSP) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("IndexToRFSP", v) }

// DecodeIndexToRFSP decodes a standalone APER value and reports the consumed octets.
func DecodeIndexToRFSP(data []byte) (IndexToRFSP, int, error) {
	var v IndexToRFSP
	err := ngapcodec.DecodeParameter("IndexToRFSP", data, &v)
	if ngapcodec.IsFatal(err) {
		return IndexToRFSP{}, 0, err
	}
	return v, len(data), err
}
