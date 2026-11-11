// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type HLComDeactivateExtIEsEntryPresent uint32

const (
	HLComDeactivateExtIEsEntryPresentNothing HLComDeactivateExtIEsEntryPresent = iota
	HLComDeactivateExtIEsEntryPresentUnknown
)

type HLComDeactivateExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      HLComDeactivateExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HLComDeactivateExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HLComDeactivate-ExtIEsEntry", v)
}

// DecodeHLComDeactivateExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeHLComDeactivateExtIEsEntry(data []byte) (HLComDeactivateExtIEsEntry, int, error) {
	var v HLComDeactivateExtIEsEntry
	err := ngapcodec.DecodeParameter("HLComDeactivate-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return HLComDeactivateExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
