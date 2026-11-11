// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceModifyItemModInd represents ASN.1 PDUSessionResourceModifyItemModInd.
type PDUSessionResourceModifyItemModInd struct {
	PDUSessionID                               PDUSessionID
	PDUSessionResourceModifyIndicationTransfer []byte
	HasIEExtensions                            bool
	IEExtensions                               []PDUSessionResourceModifyItemModIndExtIEsEntry
	UnknownExtensions                          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyItemModInd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyItemModInd", v)
}

// DecodePDUSessionResourceModifyItemModInd decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyItemModInd(data []byte) (PDUSessionResourceModifyItemModInd, int, error) {
	var v PDUSessionResourceModifyItemModInd
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyItemModInd", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyItemModInd{}, 0, err
	}
	return v, len(data), err
}
