// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UENGAPIDPair represents ASN.1 UE-NGAP-ID-pair.
type UENGAPIDPair struct {
	AMFUENGAPID       AMFUENGAPID
	RANUENGAPID       RANUENGAPID
	HasIEExtensions   bool
	IEExtensions      []UENGAPIDPairExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UENGAPIDPair) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-NGAP-ID-pair", v)
}

// DecodeUENGAPIDPair decodes a standalone APER value and reports the consumed octets.
func DecodeUENGAPIDPair(data []byte) (UENGAPIDPair, int, error) {
	var v UENGAPIDPair
	err := ngapcodec.DecodeParameter("UE-NGAP-ID-pair", data, &v)
	if ngapcodec.IsFatal(err) {
		return UENGAPIDPair{}, 0, err
	}
	return v, len(data), err
}
