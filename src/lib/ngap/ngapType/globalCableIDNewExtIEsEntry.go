// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GlobalCableIDNewExtIEsEntryPresent uint32

const (
	GlobalCableIDNewExtIEsEntryPresentNothing GlobalCableIDNewExtIEsEntryPresent = iota
	GlobalCableIDNewExtIEsEntryPresentUnknown
)

type GlobalCableIDNewExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GlobalCableIDNewExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalCableIDNewExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalCable-ID-new-ExtIEsEntry", v)
}

// DecodeGlobalCableIDNewExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalCableIDNewExtIEsEntry(data []byte) (GlobalCableIDNewExtIEsEntry, int, error) {
	var v GlobalCableIDNewExtIEsEntry
	err := ngapcodec.DecodeParameter("GlobalCable-ID-new-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalCableIDNewExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
