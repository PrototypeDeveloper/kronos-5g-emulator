// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceResumeItemRESResExtIEsEntryPresent uint32

const (
	PDUSessionResourceResumeItemRESResExtIEsEntryPresentNothing PDUSessionResourceResumeItemRESResExtIEsEntryPresent = iota
	PDUSessionResourceResumeItemRESResExtIEsEntryPresentUnknown
)

type PDUSessionResourceResumeItemRESResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceResumeItemRESResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceResumeItemRESResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceResumeItemRESRes-ExtIEsEntry", v)
}

// DecodePDUSessionResourceResumeItemRESResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceResumeItemRESResExtIEsEntry(data []byte) (PDUSessionResourceResumeItemRESResExtIEsEntry, int, error) {
	var v PDUSessionResourceResumeItemRESResExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceResumeItemRESRes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceResumeItemRESResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
