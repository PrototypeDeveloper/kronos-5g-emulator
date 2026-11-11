// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ServiceAreaInformation represents ASN.1 ServiceAreaInformation.
type ServiceAreaInformation struct {
	List []ServiceAreaInformationItem `aper:"sizeLB:1,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ServiceAreaInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ServiceAreaInformation", v)
}

// DecodeServiceAreaInformation decodes a standalone APER value and reports the consumed octets.
func DecodeServiceAreaInformation(data []byte) (ServiceAreaInformation, int, error) {
	var v ServiceAreaInformation
	err := ngapcodec.DecodeParameter("ServiceAreaInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return ServiceAreaInformation{}, 0, err
	}
	return v, len(data), err
}
