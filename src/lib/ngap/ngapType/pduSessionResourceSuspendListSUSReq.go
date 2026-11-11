// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSuspendListSUSReq represents ASN.1 PDUSessionResourceSuspendListSUSReq.
type PDUSessionResourceSuspendListSUSReq struct {
	List []PDUSessionResourceSuspendItemSUSReq `aper:"sizeLB:1,sizeUB:256"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSuspendListSUSReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSuspendListSUSReq", v)
}

// DecodePDUSessionResourceSuspendListSUSReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSuspendListSUSReq(data []byte) (PDUSessionResourceSuspendListSUSReq, int, error) {
	var v PDUSessionResourceSuspendListSUSReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceSuspendListSUSReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSuspendListSUSReq{}, 0, err
	}
	return v, len(data), err
}
