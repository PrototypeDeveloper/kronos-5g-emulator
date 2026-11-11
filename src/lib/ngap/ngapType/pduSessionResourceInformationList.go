// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceInformationList represents ASN.1 PDUSessionResourceInformationList.
type PDUSessionResourceInformationList struct {
	List []PDUSessionResourceInformationItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceInformationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceInformationList", v)
}

// DecodePDUSessionResourceInformationList decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceInformationList(data []byte) (PDUSessionResourceInformationList, int, error) {
	var v PDUSessionResourceInformationList
	err := ngapcodec.DecodeParameter("PDUSessionResourceInformationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceInformationList{}, 0, err
	}
	return v, len(data), err
}
