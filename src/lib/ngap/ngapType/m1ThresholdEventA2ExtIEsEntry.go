// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type M1ThresholdEventA2ExtIEsEntryPresent uint32

const (
	M1ThresholdEventA2ExtIEsEntryPresentNothing M1ThresholdEventA2ExtIEsEntryPresent = iota
	M1ThresholdEventA2ExtIEsEntryPresentUnknown
)

type M1ThresholdEventA2ExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      M1ThresholdEventA2ExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1ThresholdEventA2ExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1ThresholdEventA2-ExtIEsEntry", v)
}

// DecodeM1ThresholdEventA2ExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeM1ThresholdEventA2ExtIEsEntry(data []byte) (M1ThresholdEventA2ExtIEsEntry, int, error) {
	var v M1ThresholdEventA2ExtIEsEntry
	err := ngapcodec.DecodeParameter("M1ThresholdEventA2-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1ThresholdEventA2ExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
