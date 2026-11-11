// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSetupItemCxtReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceSetupItemCxtReqExtIEsEntryPresentNothing PDUSessionResourceSetupItemCxtReqExtIEsEntryPresent = iota
	PDUSessionResourceSetupItemCxtReqExtIEsEntryPresentPduSessionExpectedUEActivityBehaviour
	PDUSessionResourceSetupItemCxtReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceSetupItemCxtReqExtIEsEntry struct {
	ID                                    ProtocolExtensionID
	Criticality                           Criticality
	Present                               PDUSessionResourceSetupItemCxtReqExtIEsEntryPresent
	PduSessionExpectedUEActivityBehaviour ExpectedUEActivityBehaviour `ngap:"id:281,criticality:ignore,presence:optional"`
	UnknownValue                          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSetupItemCxtReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupItemCxtReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSetupItemCxtReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSetupItemCxtReqExtIEsEntry(data []byte) (PDUSessionResourceSetupItemCxtReqExtIEsEntry, int, error) {
	var v PDUSessionResourceSetupItemCxtReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSetupItemCxtReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSetupItemCxtReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
