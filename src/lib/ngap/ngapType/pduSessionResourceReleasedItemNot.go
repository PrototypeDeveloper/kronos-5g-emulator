// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleasedItemNot represents ASN.1 PDUSessionResourceReleasedItemNot.
type PDUSessionResourceReleasedItemNot struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceNotifyReleasedTransfer []byte
	HasIEExtensions                          bool
	IEExtensions                             []PDUSessionResourceReleasedItemNotExtIEsEntry
	UnknownExtensions                        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedItemNot) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedItemNot", v)
}

// DecodePDUSessionResourceReleasedItemNot decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedItemNot(data []byte) (PDUSessionResourceReleasedItemNot, int, error) {
	var v PDUSessionResourceReleasedItemNot
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedItemNot", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedItemNot{}, 0, err
	}
	return v, len(data), err
}
