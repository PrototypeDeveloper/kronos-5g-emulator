// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo represents ASN.1 NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo.
const (
	NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfoPresentPresent    Enumerated = 0
	NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfoPresentNotPresent Enumerated = 1
)

type NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo", v)
}

// DecodeNGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo decodes a standalone APER value and reports consumed octets.
func DecodeNGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo(data []byte) (NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo, int, error) {
	var v NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo
	err := ngapcodec.DecodeParameter("inline:NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfo{}, 0, err
	}
	return v, len(data), err
}
