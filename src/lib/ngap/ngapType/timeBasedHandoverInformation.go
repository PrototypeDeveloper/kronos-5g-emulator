// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeBasedHandoverInformation represents ASN.1 TimeBasedHandoverInformation.
type TimeBasedHandoverInformation struct {
	HOWindowStart     HandoverWindowStart
	HOWindowDuration  HandoverWindowDuration
	HasIEExtensions   bool
	IEExtensions      []TimeBasedHandoverInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeBasedHandoverInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TimeBasedHandoverInformation", v)
}

// DecodeTimeBasedHandoverInformation decodes a standalone APER value and reports the consumed octets.
func DecodeTimeBasedHandoverInformation(data []byte) (TimeBasedHandoverInformation, int, error) {
	var v TimeBasedHandoverInformation
	err := ngapcodec.DecodeParameter("TimeBasedHandoverInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeBasedHandoverInformation{}, 0, err
	}
	return v, len(data), err
}
