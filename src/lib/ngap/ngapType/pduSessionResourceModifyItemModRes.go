// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyItemModRes represents ASN.1 PDUSessionResourceModifyItemModRes.
type PDUSessionResourceModifyItemModRes struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceModifyResponseTransfer []byte
	HasIEExtensions                          bool
	IEExtensions                             []PDUSessionResourceModifyItemModResExtIEsEntry
	UnknownExtensions                        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyItemModRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyItemModRes", v)
}

// DecodePDUSessionResourceModifyItemModRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyItemModRes(data []byte) (PDUSessionResourceModifyItemModRes, int, error) {
	var v PDUSessionResourceModifyItemModRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyItemModRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyItemModRes{}, 0, err
	}
	return v, len(data), err
}
