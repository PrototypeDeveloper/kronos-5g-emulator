// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceFailedToModifyItemModCfmExtIEsEntryPresent uint32

const (
	PDUSessionResourceFailedToModifyItemModCfmExtIEsEntryPresentNothing PDUSessionResourceFailedToModifyItemModCfmExtIEsEntryPresent = iota
	PDUSessionResourceFailedToModifyItemModCfmExtIEsEntryPresentUnknown
)

type PDUSessionResourceFailedToModifyItemModCfmExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceFailedToModifyItemModCfmExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceFailedToModifyItemModCfmExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceFailedToModifyItemModCfm-ExtIEsEntry", v)
}

// DecodePDUSessionResourceFailedToModifyItemModCfmExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceFailedToModifyItemModCfmExtIEsEntry(data []byte) (PDUSessionResourceFailedToModifyItemModCfmExtIEsEntry, int, error) {
	var v PDUSessionResourceFailedToModifyItemModCfmExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceFailedToModifyItemModCfm-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceFailedToModifyItemModCfmExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
