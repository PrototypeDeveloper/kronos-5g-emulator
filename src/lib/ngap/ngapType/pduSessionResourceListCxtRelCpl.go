// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceListCxtRelCpl represents ASN.1 PDUSessionResourceListCxtRelCpl.
type PDUSessionResourceListCxtRelCpl struct {
	List []PDUSessionResourceItemCxtRelCpl `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceListCxtRelCpl) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceListCxtRelCpl", v)
}

// DecodePDUSessionResourceListCxtRelCpl decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceListCxtRelCpl(data []byte) (PDUSessionResourceListCxtRelCpl, int, error) {
	var v PDUSessionResourceListCxtRelCpl
	err := ngapcodec.DecodeParameter("PDUSessionResourceListCxtRelCpl", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceListCxtRelCpl{}, 0, err
	}
	return v, len(data), err
}
