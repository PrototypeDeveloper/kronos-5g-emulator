// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceHandoverItemExtIEsEntryPresent uint32

const (
	PDUSessionResourceHandoverItemExtIEsEntryPresentNothing PDUSessionResourceHandoverItemExtIEsEntryPresent = iota
	PDUSessionResourceHandoverItemExtIEsEntryPresentUnknown
)

type PDUSessionResourceHandoverItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceHandoverItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceHandoverItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceHandoverItem-ExtIEsEntry", v)
}

// DecodePDUSessionResourceHandoverItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceHandoverItemExtIEsEntry(data []byte) (PDUSessionResourceHandoverItemExtIEsEntry, int, error) {
	var v PDUSessionResourceHandoverItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceHandoverItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceHandoverItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
