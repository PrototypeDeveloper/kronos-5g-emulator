// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANNoofRRCConnections represents ASN.1 NGRAN-NoofRRCConnections.
type NGRANNoofRRCConnections struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:65536"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANNoofRRCConnections) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-NoofRRCConnections", v)
}

// DecodeNGRANNoofRRCConnections decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANNoofRRCConnections(data []byte) (NGRANNoofRRCConnections, int, error) {
	var v NGRANNoofRRCConnections
	err := ngapcodec.DecodeParameter("NGRAN-NoofRRCConnections", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANNoofRRCConnections{}, 0, err
	}
	return v, len(data), err
}
