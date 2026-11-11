// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSetupItemSUReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceSetupItemSUReqExtIEsEntryPresentNothing PDUSessionResourceSetupItemSUReqExtIEsEntryPresent = iota
	PDUSessionResourceSetupItemSUReqExtIEsEntryPresentPduSessionExpectedUEActivityBehaviour
	PDUSessionResourceSetupItemSUReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceSetupItemSUReqExtIEsEntry struct {
	ID                                    ProtocolExtensionID
	Criticality                           Criticality
	Present                               PDUSessionResourceSetupItemSUReqExtIEsEntryPresent
	PduSessionExpectedUEActivityBehaviour ExpectedUEActivityBehaviour `ngap:"id:281,criticality:ignore,presence:optional"`
	UnknownValue                          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemSUReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemSUReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSetupItemSUReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemSUReqExtIEsEntry(data []byte) (PDUSessionResourceSetupItemSUReqExtIEsEntry, int, error) {
	var v PDUSessionResourceSetupItemSUReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemSUReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemSUReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
