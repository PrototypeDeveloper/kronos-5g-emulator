// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceToBeSwitchedDLItem represents ASN.1 PDUSessionResourceToBeSwitchedDLItem.
type PDUSessionResourceToBeSwitchedDLItem struct {
	PDUSessionID              PDUSessionID
	PathSwitchRequestTransfer []byte
	HasIEExtensions           bool
	IEExtensions              []PDUSessionResourceToBeSwitchedDLItemExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToBeSwitchedDLItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToBeSwitchedDLItem", v)
}

// DecodePDUSessionResourceToBeSwitchedDLItem decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToBeSwitchedDLItem(data []byte) (PDUSessionResourceToBeSwitchedDLItem, int, error) {
	var v PDUSessionResourceToBeSwitchedDLItem
	err := ngapcodec.DecodeParameter("PDUSessionResourceToBeSwitchedDLItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToBeSwitchedDLItem{}, 0, err
	}
	return v, len(data), err
}
