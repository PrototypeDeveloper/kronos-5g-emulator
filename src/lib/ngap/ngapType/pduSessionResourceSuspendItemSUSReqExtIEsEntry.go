// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSuspendItemSUSReqExtIEsEntryPresent uint32

const (
	PDUSessionResourceSuspendItemSUSReqExtIEsEntryPresentNothing PDUSessionResourceSuspendItemSUSReqExtIEsEntryPresent = iota
	PDUSessionResourceSuspendItemSUSReqExtIEsEntryPresentUnknown
)

type PDUSessionResourceSuspendItemSUSReqExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceSuspendItemSUSReqExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSuspendItemSUSReqExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSuspendItemSUSReq-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSuspendItemSUSReqExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSuspendItemSUSReqExtIEsEntry(data []byte) (PDUSessionResourceSuspendItemSUSReqExtIEsEntry, int, error) {
	var v PDUSessionResourceSuspendItemSUSReqExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSuspendItemSUSReq-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSuspendItemSUSReqExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
