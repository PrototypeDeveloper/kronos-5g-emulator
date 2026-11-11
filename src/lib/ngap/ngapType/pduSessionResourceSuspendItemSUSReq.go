// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSuspendItemSUSReq represents ASN.1 PDUSessionResourceSuspendItemSUSReq.
type PDUSessionResourceSuspendItemSUSReq struct {
	PDUSessionID                    PDUSessionID
	UEContextSuspendRequestTransfer []byte
	HasIEExtensions                 bool
	IEExtensions                    []PDUSessionResourceSuspendItemSUSReqExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSuspendItemSUSReq) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSuspendItemSUSReq", v)
}

// DecodePDUSessionResourceSuspendItemSUSReq decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSuspendItemSUSReq(data []byte) (PDUSessionResourceSuspendItemSUSReq, int, error) {
	var v PDUSessionResourceSuspendItemSUSReq
	err := ngapcodec.DecodeParameter("PDUSessionResourceSuspendItemSUSReq", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSuspendItemSUSReq{}, 0, err
	}
	return v, len(data), err
}
