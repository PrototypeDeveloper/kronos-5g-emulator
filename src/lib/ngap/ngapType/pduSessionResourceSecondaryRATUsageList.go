// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSecondaryRATUsageList represents ASN.1 PDUSessionResourceSecondaryRATUsageList.
type PDUSessionResourceSecondaryRATUsageList struct {
	List []PDUSessionResourceSecondaryRATUsageItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSecondaryRATUsageList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSecondaryRATUsageList", v)
}

// DecodePDUSessionResourceSecondaryRATUsageList decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSecondaryRATUsageList(data []byte) (PDUSessionResourceSecondaryRATUsageList, int, error) {
	var v PDUSessionResourceSecondaryRATUsageList
	err := ngapcodec.DecodeParameter("PDUSessionResourceSecondaryRATUsageList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSecondaryRATUsageList{}, 0, err
	}
	return v, len(data), err
}
