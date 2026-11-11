// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo represents ASN.1 NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo.
const (
	NGAPIESupportInformationResponseItemNgapProtocolIESupportInfoPresentSupported    Enumerated = 0
	NGAPIESupportInformationResponseItemNgapProtocolIESupportInfoPresentNotSupported Enumerated = 1
)

type NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo", v)
}

// DecodeNGAPIESupportInformationResponseItemNgapProtocolIESupportInfo decodes a standalone APER value and reports consumed octets.
func DecodeNGAPIESupportInformationResponseItemNgapProtocolIESupportInfo(data []byte) (NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo, int, error) {
	var v NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo
	err := ngapcodec.DecodeParameter("inline:NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPIESupportInformationResponseItemNgapProtocolIESupportInfo{}, 0, err
	}
	return v, len(data), err
}
