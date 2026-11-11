// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceSwitchedItemExtIEsEntryPresent uint32

const (
	PDUSessionResourceSwitchedItemExtIEsEntryPresentNothing PDUSessionResourceSwitchedItemExtIEsEntryPresent = iota
	PDUSessionResourceSwitchedItemExtIEsEntryPresentPduSessionExpectedUEActivityBehaviour
	PDUSessionResourceSwitchedItemExtIEsEntryPresentUnknown
)

type PDUSessionResourceSwitchedItemExtIEsEntry struct {
	ID                                    ProtocolExtensionID
	Criticality                           Criticality
	Present                               PDUSessionResourceSwitchedItemExtIEsEntryPresent
	PduSessionExpectedUEActivityBehaviour ExpectedUEActivityBehaviour `ngap:"id:281,criticality:ignore,presence:optional"`
	UnknownValue                          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceSwitchedItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceSwitchedItem-ExtIEsEntry", v)
}

// DecodePDUSessionResourceSwitchedItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceSwitchedItemExtIEsEntry(data []byte) (PDUSessionResourceSwitchedItemExtIEsEntry, int, error) {
	var v PDUSessionResourceSwitchedItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceSwitchedItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceSwitchedItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
