// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToModifyItemModCfm represents ASN.1 PDUSessionResourceFailedToModifyItemModCfm.
type PDUSessionResourceFailedToModifyItemModCfm struct {
	PDUSessionID                                           PDUSessionID
	PDUSessionResourceModifyIndicationUnsuccessfulTransfer []byte
	HasIEExtensions                                        bool
	IEExtensions                                           []PDUSessionResourceFailedToModifyItemModCfmExtIEsEntry
	UnknownExtensions                                      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToModifyItemModCfm) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToModifyItemModCfm", v)
}

// DecodePDUSessionResourceFailedToModifyItemModCfm decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToModifyItemModCfm(data []byte) (PDUSessionResourceFailedToModifyItemModCfm, int, error) {
	var v PDUSessionResourceFailedToModifyItemModCfm
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToModifyItemModCfm", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToModifyItemModCfm{}, 0, err
	}
	return v, len(data), err
}
