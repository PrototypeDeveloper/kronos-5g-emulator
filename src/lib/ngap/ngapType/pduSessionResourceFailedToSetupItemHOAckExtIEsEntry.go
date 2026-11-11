// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToSetupItemHOAckExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToSetupItemHOAckExtIEsEntryPresentNothing PDUSessionResourceFailedToSetupItemHOAckExtIEsEntryPresent = iota
	PDUSessionResourceFailedToSetupItemHOAckExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToSetupItemHOAckExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToSetupItemHOAckExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemHOAckExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemHOAck-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToSetupItemHOAckExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemHOAckExtIEsEntry(data []byte) (PDUSessionResourceFailedToSetupItemHOAckExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToSetupItemHOAckExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemHOAck-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemHOAckExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
