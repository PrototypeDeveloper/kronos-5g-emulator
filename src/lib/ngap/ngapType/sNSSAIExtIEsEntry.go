// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SNSSAIExtIEsEntryPresent uint32

const (
	SNSSAIExtIEsEntryPresentNothing SNSSAIExtIEsEntryPresent = iota
	SNSSAIExtIEsEntryPresentUnknown
)

type SNSSAIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SNSSAIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNSSAIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("S-NSSAI-ExtIEsEntry", v)
}

// DecodeSNSSAIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSNSSAIExtIEsEntry(data []byte) (SNSSAIExtIEsEntry, int, error) {
	var v SNSSAIExtIEsEntry
	err := ngapcodec.DecodeParameter("S-NSSAI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNSSAIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
