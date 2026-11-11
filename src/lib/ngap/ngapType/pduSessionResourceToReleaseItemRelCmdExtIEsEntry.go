// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceToReleaseItemRelCmdExtIEsEntryPresent uint32

const (
	PDUSessionResourceToReleaseItemRelCmdExtIEsEntryPresentNothing PDUSessionResourceToReleaseItemRelCmdExtIEsEntryPresent = iota
	PDUSessionResourceToReleaseItemRelCmdExtIEsEntryPresentUnknown
)

type PDUSessionResourceToReleaseItemRelCmdExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceToReleaseItemRelCmdExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceToReleaseItemRelCmdExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceToReleaseItemRelCmd-ExtIEsEntry", v)
}

// DecodePDUSessionResourceToReleaseItemRelCmdExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceToReleaseItemRelCmdExtIEsEntry(data []byte) (PDUSessionResourceToReleaseItemRelCmdExtIEsEntry, int, error) {
	var v PDUSessionResourceToReleaseItemRelCmdExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceToReleaseItemRelCmd-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceToReleaseItemRelCmdExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
