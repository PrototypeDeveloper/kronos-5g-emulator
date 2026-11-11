// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceToReleaseItemHOCmdExtIEsEntryPresent uint32

const (
	PDUSessionResourceToReleaseItemHOCmdExtIEsEntryPresentNothing PDUSessionResourceToReleaseItemHOCmdExtIEsEntryPresent = iota
	PDUSessionResourceToReleaseItemHOCmdExtIEsEntryPresentUnknown
)

type PDUSessionResourceToReleaseItemHOCmdExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceToReleaseItemHOCmdExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToReleaseItemHOCmdExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToReleaseItemHOCmd-ExtIEsEntry", v)
}

// DecodePDUSessionResourceToReleaseItemHOCmdExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToReleaseItemHOCmdExtIEsEntry(data []byte) (PDUSessionResourceToReleaseItemHOCmdExtIEsEntry, int, error) {
	var v PDUSessionResourceToReleaseItemHOCmdExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceToReleaseItemHOCmd-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToReleaseItemHOCmdExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
