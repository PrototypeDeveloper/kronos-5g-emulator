// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANTraceID represents ASN.1 NGRANTraceID.
type NGRANTraceID struct {
	Value []byte `aper:"sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANTraceID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NGRANTraceID", v) }

// DecodeNGRANTraceID decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANTraceID(data []byte) (NGRANTraceID, int, error) {
	var v NGRANTraceID
	err := ngapcodec.DecodeParameter("NGRANTraceID", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANTraceID{}, 0, err
	}
	return v, len(data), err
}
