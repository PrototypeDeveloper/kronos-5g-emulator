// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceItemHORqdExtIEsEntryPresent uint32

const (
	PDUSessionResourceItemHORqdExtIEsEntryPresentNothing PDUSessionResourceItemHORqdExtIEsEntryPresent = iota
	PDUSessionResourceItemHORqdExtIEsEntryPresentUnknown
)

type PDUSessionResourceItemHORqdExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceItemHORqdExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceItemHORqdExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceItemHORqd-ExtIEsEntry", v)
}

// DecodePDUSessionResourceItemHORqdExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceItemHORqdExtIEsEntry(data []byte) (PDUSessionResourceItemHORqdExtIEsEntry, int, error) {
	var v PDUSessionResourceItemHORqdExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceItemHORqd-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceItemHORqdExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
