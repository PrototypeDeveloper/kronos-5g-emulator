// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NTNGeographicalAreaList represents ASN.1 NTN-Geographical-Area-List.
type NTNGeographicalAreaList struct {
	List []NTNGeographicalAreaItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NTNGeographicalAreaList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NTN-Geographical-Area-List", v)
}

// DecodeNTNGeographicalAreaList decodes a standalone APER value and reports the consumed octets.
func DecodeNTNGeographicalAreaList(data []byte) (NTNGeographicalAreaList, int, error) {
	var v NTNGeographicalAreaList
	err := ngapcodec.DecodeParameter("NTN-Geographical-Area-List", data, &v)
	if ngapcodec.IsFatal(err) {
		return NTNGeographicalAreaList{}, 0, err
	}
	return v, len(data), err
}
