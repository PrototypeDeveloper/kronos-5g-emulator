// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type HFCNodeIDNewExtIEsEntryPresent uint32

const (
	HFCNodeIDNewExtIEsEntryPresentNothing HFCNodeIDNewExtIEsEntryPresent = iota
	HFCNodeIDNewExtIEsEntryPresentGlobalCableID
	HFCNodeIDNewExtIEsEntryPresentUnknown
)

type HFCNodeIDNewExtIEsEntry struct {
	ID            ProtocolExtensionID
	Criticality   Criticality
	Present       HFCNodeIDNewExtIEsEntryPresent
	GlobalCableID GlobalCableID `ngap:"id:275,criticality:ignore,presence:optional"`
	UnknownValue  []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HFCNodeIDNewExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HFCNode-ID-new-ExtIEsEntry", v)
}

// DecodeHFCNodeIDNewExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeHFCNodeIDNewExtIEsEntry(data []byte) (HFCNodeIDNewExtIEsEntry, int, error) {
	var v HFCNodeIDNewExtIEsEntry
	err := ngapcodec.DecodeParameter("HFCNode-ID-new-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return HFCNodeIDNewExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
