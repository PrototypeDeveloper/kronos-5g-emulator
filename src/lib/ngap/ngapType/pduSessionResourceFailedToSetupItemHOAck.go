// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionResourceFailedToSetupItemHOAck represents ASN.1 PDUSessionResourceFailedToSetupItemHOAck.
type PDUSessionResourceFailedToSetupItemHOAck struct {
	PDUSessionID                                   PDUSessionID
	HandoverResourceAllocationUnsuccessfulTransfer []byte
	HasIEExtensions                                bool
	IEExtensions                                   []PDUSessionResourceFailedToSetupItemHOAckExtIEsEntry
	UnknownExtensions                              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemHOAck) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemHOAck", v)
}

// DecodePDUSessionResourceFailedToSetupItemHOAck decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemHOAck(data []byte) (PDUSessionResourceFailedToSetupItemHOAck, int, error) {
	var v PDUSessionResourceFailedToSetupItemHOAck
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemHOAck", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemHOAck{}, 0, err
	}
	return v, len(data), err
}
