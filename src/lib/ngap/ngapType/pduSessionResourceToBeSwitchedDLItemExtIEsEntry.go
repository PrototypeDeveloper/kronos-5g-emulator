// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceToBeSwitchedDLItemExtIEsEntryPresent uint32

const (
	PDUSessionResourceToBeSwitchedDLItemExtIEsEntryPresentNothing PDUSessionResourceToBeSwitchedDLItemExtIEsEntryPresent = iota
	PDUSessionResourceToBeSwitchedDLItemExtIEsEntryPresentUnknown
)

type PDUSessionResourceToBeSwitchedDLItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceToBeSwitchedDLItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToBeSwitchedDLItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToBeSwitchedDLItem-ExtIEsEntry", v)
}

// DecodePDUSessionResourceToBeSwitchedDLItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToBeSwitchedDLItemExtIEsEntry(data []byte) (PDUSessionResourceToBeSwitchedDLItemExtIEsEntry, int, error) {
	var v PDUSessionResourceToBeSwitchedDLItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceToBeSwitchedDLItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToBeSwitchedDLItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
