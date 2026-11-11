// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyItemModReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyItemModReqExtIEsEntryPresentNothing PDUSessionResourceModifyItemModReqExtIEsEntryPresent = iota
	PDUSessionResourceModifyItemModReqExtIEsEntryPresentSNSSAI
	PDUSessionResourceModifyItemModReqExtIEsEntryPresentPduSessionExpectedUEActivityBehaviour
	PDUSessionResourceModifyItemModReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyItemModReqExtIEsEntry struct {
	ID                                    ProtocolExtensionID
	Criticality                           Criticality
	Present                               PDUSessionResourceModifyItemModReqExtIEsEntryPresent
	SNSSAI                                SNSSAI                      `ngap:"id:148,criticality:reject,presence:optional"`
	PduSessionExpectedUEActivityBehaviour ExpectedUEActivityBehaviour `ngap:"id:281,criticality:ignore,presence:optional"`
	UnknownValue                          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyItemModReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyItemModReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyItemModReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyItemModReqExtIEsEntry(data []byte) (PDUSessionResourceModifyItemModReqExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyItemModReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyItemModReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyItemModReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
