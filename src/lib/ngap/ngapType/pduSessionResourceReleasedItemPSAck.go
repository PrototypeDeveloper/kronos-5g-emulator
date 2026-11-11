// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceReleasedItemPSAck represents ASN.1 PDUSessionResourceReleasedItemPSAck.
type PDUSessionResourceReleasedItemPSAck struct {
	PDUSessionID                          PDUSessionID
	PathSwitchRequestUnsuccessfulTransfer []byte
	HasIEExtensions                       bool
	IEExtensions                          []PDUSessionResourceReleasedItemPSAckExtIEsEntry
	UnknownExtensions                     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceReleasedItemPSAck) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceReleasedItemPSAck", v)
}

// DecodePDUSessionResourceReleasedItemPSAck decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceReleasedItemPSAck(data []byte) (PDUSessionResourceReleasedItemPSAck, int, error) {
	var v PDUSessionResourceReleasedItemPSAck
	err := ngapcodec.DecodeParameter("PDUSessionResourceReleasedItemPSAck", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceReleasedItemPSAck{}, 0, err
	}
	return v, len(data), err
}
