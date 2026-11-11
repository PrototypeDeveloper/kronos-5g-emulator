// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type HLComActivateExtIEsEntryPresent uint32

const (
	HLComActivateExtIEsEntryPresentNothing HLComActivateExtIEsEntryPresent = iota
	HLComActivateExtIEsEntryPresentUnknown
)

type HLComActivateExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      HLComActivateExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HLComActivateExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HLComActivate-ExtIEsEntry", v)
}

// DecodeHLComActivateExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeHLComActivateExtIEsEntry(data []byte) (HLComActivateExtIEsEntry, int, error) {
	var v HLComActivateExtIEsEntry
	err := ngapcodec.DecodeParameter("HLComActivate-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return HLComActivateExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
