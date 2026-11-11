// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceModifyItemModCfmExtIEsEntryPresent uint32

const (
	PDUSessionResourceModifyItemModCfmExtIEsEntryPresentNothing PDUSessionResourceModifyItemModCfmExtIEsEntryPresent = iota
	PDUSessionResourceModifyItemModCfmExtIEsEntryPresentUnknown
)

type PDUSessionResourceModifyItemModCfmExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionResourceModifyItemModCfmExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceModifyItemModCfmExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceModifyItemModCfm-ExtIEsEntry", v)
}

// DecodePDUSessionResourceModifyItemModCfmExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceModifyItemModCfmExtIEsEntry(data []byte) (PDUSessionResourceModifyItemModCfmExtIEsEntry, int, error) {
	var v PDUSessionResourceModifyItemModCfmExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceModifyItemModCfm-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceModifyItemModCfmExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
