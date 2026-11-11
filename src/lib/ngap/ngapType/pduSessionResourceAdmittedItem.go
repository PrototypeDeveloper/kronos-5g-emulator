// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceAdmittedItem represents ASN.1 PDUSessionResourceAdmittedItem.
type PDUSessionResourceAdmittedItem struct {
	PDUSessionID                       PDUSessionID
	HandoverRequestAcknowledgeTransfer []byte
	HasIEExtensions                    bool
	IEExtensions                       []PDUSessionResourceAdmittedItemExtIEsEntry
	UnknownExtensions                  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceAdmittedItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceAdmittedItem", v)
}

// DecodePDUSessionResourceAdmittedItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceAdmittedItem(data []byte) (PDUSessionResourceAdmittedItem, int, error) {
	var v PDUSessionResourceAdmittedItem
	err := ngapcodec.DecodeParameter("PDUSessionResourceAdmittedItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceAdmittedItem{}, 0, err
	}
	return v, len(data), err
}
