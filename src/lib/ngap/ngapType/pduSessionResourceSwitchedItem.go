// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceSwitchedItem represents ASN.1 PDUSessionResourceSwitchedItem.
type PDUSessionResourceSwitchedItem struct {
	PDUSessionID                         PDUSessionID
	PathSwitchRequestAcknowledgeTransfer []byte
	HasIEExtensions                      bool
	IEExtensions                         []PDUSessionResourceSwitchedItemExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSwitchedItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSwitchedItem", v)
}

// DecodePDUSessionResourceSwitchedItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSwitchedItem(data []byte) (PDUSessionResourceSwitchedItem, int, error) {
	var v PDUSessionResourceSwitchedItem
	err := ngapcodec.DecodeParameter("PDUSessionResourceSwitchedItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSwitchedItem{}, 0, err
	}
	return v, len(data), err
}
