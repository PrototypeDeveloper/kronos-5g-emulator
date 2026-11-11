// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSetupItemHOReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceSetupItemHOReqExtIEsEntryPresentNothing PDUSessionResourceSetupItemHOReqExtIEsEntryPresent = iota
	PDUSessionResourceSetupItemHOReqExtIEsEntryPresentPduSessionExpectedUEActivityBehaviour
	PDUSessionResourceSetupItemHOReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceSetupItemHOReqExtIEsEntry struct {
	ID                                    ProtocolExtensionID
	Criticality                           Criticality
	Present                               PDUSessionResourceSetupItemHOReqExtIEsEntryPresent
	PduSessionExpectedUEActivityBehaviour ExpectedUEActivityBehaviour `ngap:"id:281,criticality:ignore,presence:optional"`
	UnknownValue                          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemHOReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemHOReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSetupItemHOReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemHOReqExtIEsEntry(data []byte) (PDUSessionResourceSetupItemHOReqExtIEsEntry, int, error) {
	var v PDUSessionResourceSetupItemHOReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemHOReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemHOReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
