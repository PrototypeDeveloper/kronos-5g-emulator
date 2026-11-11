// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceInformationItemExtIEsEntryPresent uint32

const (
	PDUSessionResourceInformationItemExtIEsEntryPresentNothing PDUSessionResourceInformationItemExtIEsEntryPresent = iota
	PDUSessionResourceInformationItemExtIEsEntryPresentUnknown
)

type PDUSessionResourceInformationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceInformationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceInformationItem-ExtIEsEntry", v)
}

// DecodePDUSessionResourceInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceInformationItemExtIEsEntry(data []byte) (PDUSessionResourceInformationItemExtIEsEntry, int, error) {
	var v PDUSessionResourceInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceInformationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
