// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToSetupItemPSReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToSetupItemPSReqExtIEsEntryPresentNothing PDUSessionResourceFailedToSetupItemPSReqExtIEsEntryPresent = iota
	PDUSessionResourceFailedToSetupItemPSReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToSetupItemPSReqExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToSetupItemPSReqExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToSetupItemPSReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToSetupItemPSReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToSetupItemPSReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToSetupItemPSReqExtIEsEntry(data []byte) (PDUSessionResourceFailedToSetupItemPSReqExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToSetupItemPSReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToSetupItemPSReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToSetupItemPSReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
