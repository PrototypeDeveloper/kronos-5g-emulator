// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionForPagingList represents ASN.1 PDUSessionForPagingList.
type PDUSessionForPagingList struct {
	List []PDUSessionForPagingItem `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionForPagingList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionForPagingList", v)
}

// DecodePDUSessionForPagingList decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionForPagingList(data []byte) (PDUSessionForPagingList, int, error) {
	var v PDUSessionForPagingList
	err := ngapcodec.DecodeParameter("PDUSessionForPagingList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionForPagingList{}, 0, err
	}
	return v, len(data), err
}
