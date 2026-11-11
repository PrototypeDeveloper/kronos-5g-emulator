// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyItemModCfm represents ASN.1 PDUSessionResourceModifyItemModCfm.
type PDUSessionResourceModifyItemModCfm struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceModifyConfirmTransfer []byte
	HasIEExtensions                         bool
	IEExtensions                            []PDUSessionResourceModifyItemModCfmExtIEsEntry
	UnknownExtensions                       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyItemModCfm) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyItemModCfm", v)
}

// DecodePDUSessionResourceModifyItemModCfm decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyItemModCfm(data []byte) (PDUSessionResourceModifyItemModCfm, int, error) {
	var v PDUSessionResourceModifyItemModCfm
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyItemModCfm", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyItemModCfm{}, 0, err
	}
	return v, len(data), err
}
