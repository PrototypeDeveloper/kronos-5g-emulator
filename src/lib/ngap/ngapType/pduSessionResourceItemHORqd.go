// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceItemHORqd represents ASN.1 PDUSessionResourceItemHORqd.
type PDUSessionResourceItemHORqd struct {
	PDUSessionID             PDUSessionID
	HandoverRequiredTransfer []byte
	HasIEExtensions          bool
	IEExtensions             []PDUSessionResourceItemHORqdExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceItemHORqd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceItemHORqd", v)
}

// DecodePDUSessionResourceItemHORqd decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceItemHORqd(data []byte) (PDUSessionResourceItemHORqd, int, error) {
	var v PDUSessionResourceItemHORqd
	err := ngapcodec.DecodeParameter("PDUSessionResourceItemHORqd", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceItemHORqd{}, 0, err
	}
	return v, len(data), err
}
