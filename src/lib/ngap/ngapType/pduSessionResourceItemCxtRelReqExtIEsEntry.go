// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceItemCxtRelReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceItemCxtRelReqExtIEsEntryPresentNothing PDUSessionResourceItemCxtRelReqExtIEsEntryPresent = iota
	PDUSessionResourceItemCxtRelReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceItemCxtRelReqExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceItemCxtRelReqExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceItemCxtRelReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceItemCxtRelReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceItemCxtRelReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceItemCxtRelReqExtIEsEntry(data []byte) (PDUSessionResourceItemCxtRelReqExtIEsEntry, int, error) {
	var v PDUSessionResourceItemCxtRelReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceItemCxtRelReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceItemCxtRelReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
