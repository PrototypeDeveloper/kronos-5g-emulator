// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFUENGAPID represents ASN.1 AMF-UE-NGAP-ID.
type AMFUENGAPID struct {
	Value int64 `aper:"valueLB:0,valueUB:1099511627775"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFUENGAPID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("AMF-UE-NGAP-ID", v) }

// DecodeAMFUENGAPID decodes a standalone APER value and reports the consumed octets.
func DecodeAMFUENGAPID(data []byte) (AMFUENGAPID, int, error) {
	var v AMFUENGAPID
	err := ngapcodec.DecodeParameter("AMF-UE-NGAP-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFUENGAPID{}, 0, err
	}
	return v, len(data), err
}
