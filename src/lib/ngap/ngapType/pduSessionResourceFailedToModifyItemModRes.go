// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToModifyItemModRes represents ASN.1 PDUSessionResourceFailedToModifyItemModRes.
type PDUSessionResourceFailedToModifyItemModRes struct {
	PDUSessionID                                 PDUSessionID
	PDUSessionResourceModifyUnsuccessfulTransfer []byte
	HasIEExtensions                              bool
	IEExtensions                                 []PDUSessionResourceFailedToModifyItemModResExtIEsEntry
	UnknownExtensions                            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToModifyItemModRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToModifyItemModRes", v)
}

// DecodePDUSessionResourceFailedToModifyItemModRes decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToModifyItemModRes(data []byte) (PDUSessionResourceFailedToModifyItemModRes, int, error) {
	var v PDUSessionResourceFailedToModifyItemModRes
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToModifyItemModRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToModifyItemModRes{}, 0, err
	}
	return v, len(data), err
}
