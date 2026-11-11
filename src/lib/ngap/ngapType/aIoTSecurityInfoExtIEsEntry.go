// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AIoTSecurityInfoExtIEsEntryPresent uint32

const (
	AIoTSecurityInfoExtIEsEntryPresentNothing AIoTSecurityInfoExtIEsEntryPresent = iota
	AIoTSecurityInfoExtIEsEntryPresentUnknown
)

type AIoTSecurityInfoExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      AIoTSecurityInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIoTSecurityInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIoT-SecurityInfo-ExtIEsEntry", v)
}

// DecodeAIoTSecurityInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAIoTSecurityInfoExtIEsEntry(data []byte) (AIoTSecurityInfoExtIEsEntry, int, error) {
	var v AIoTSecurityInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("AIoT-SecurityInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIoTSecurityInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
