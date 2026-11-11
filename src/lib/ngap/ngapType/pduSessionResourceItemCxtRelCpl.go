// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceItemCxtRelCpl represents ASN.1 PDUSessionResourceItemCxtRelCpl.
type PDUSessionResourceItemCxtRelCpl struct {
	PDUSessionID      PDUSessionID
	HasIEExtensions   bool
	IEExtensions      []PDUSessionResourceItemCxtRelCplExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceItemCxtRelCpl) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceItemCxtRelCpl", v)
}

// DecodePDUSessionResourceItemCxtRelCpl decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceItemCxtRelCpl(data []byte) (PDUSessionResourceItemCxtRelCpl, int, error) {
	var v PDUSessionResourceItemCxtRelCpl
	err := ngapcodec.DecodeParameter("PDUSessionResourceItemCxtRelCpl", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceItemCxtRelCpl{}, 0, err
	}
	return v, len(data), err
}
