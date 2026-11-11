// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceListHORqd represents ASN.1 PDUSessionResourceListHORqd.
type PDUSessionResourceListHORqd struct {
	List []PDUSessionResourceItemHORqd `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceListHORqd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceListHORqd", v)
}

// DecodePDUSessionResourceListHORqd decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceListHORqd(data []byte) (PDUSessionResourceListHORqd, int, error) {
	var v PDUSessionResourceListHORqd
	err := ngapcodec.DecodeParameter("PDUSessionResourceListHORqd", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceListHORqd{}, 0, err
	}
	return v, len(data), err
}
