// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FiveGSTMSIExtIEsEntryPresent uint32

const (
	FiveGSTMSIExtIEsEntryPresentNothing FiveGSTMSIExtIEsEntryPresent = iota
	FiveGSTMSIExtIEsEntryPresentUnknown
)

type FiveGSTMSIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      FiveGSTMSIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGSTMSIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveG-S-TMSI-ExtIEsEntry", v)
}

// DecodeFiveGSTMSIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGSTMSIExtIEsEntry(data []byte) (FiveGSTMSIExtIEsEntry, int, error) {
	var v FiveGSTMSIExtIEsEntry
	err := ngapcodec.DecodeParameter("FiveG-S-TMSI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGSTMSIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
