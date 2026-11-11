// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceItemCxtRelCplExtIEsEntryPresent uint32

const (
	PDUSessionResourceItemCxtRelCplExtIEsEntryPresentNothing PDUSessionResourceItemCxtRelCplExtIEsEntryPresent = iota
	PDUSessionResourceItemCxtRelCplExtIEsEntryPresentPDUSessionResourceReleaseResponseTransfer
	PDUSessionResourceItemCxtRelCplExtIEsEntryPresentUnknown
)

type PDUSessionResourceItemCxtRelCplExtIEsEntry struct {
	ID                                        ProtocolExtensionID
	Criticality                               Criticality
	Present                                   PDUSessionResourceItemCxtRelCplExtIEsEntryPresent
	PDUSessionResourceReleaseResponseTransfer PDUSessionResourceReleaseResponseTransferOCTETSTRING `ngap:"id:145,criticality:ignore,presence:optional"`
	UnknownValue                              []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceItemCxtRelCplExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceItemCxtRelCpl-ExtIEsEntry", v)
}

// DecodePDUSessionResourceItemCxtRelCplExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceItemCxtRelCplExtIEsEntry(data []byte) (PDUSessionResourceItemCxtRelCplExtIEsEntry, int, error) {
	var v PDUSessionResourceItemCxtRelCplExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceItemCxtRelCpl-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceItemCxtRelCplExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
