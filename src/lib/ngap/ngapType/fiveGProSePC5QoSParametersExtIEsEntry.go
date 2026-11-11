// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FiveGProSePC5QoSParametersExtIEsEntryPresent uint32

const (
	FiveGProSePC5QoSParametersExtIEsEntryPresentNothing FiveGProSePC5QoSParametersExtIEsEntryPresent = iota
	FiveGProSePC5QoSParametersExtIEsEntryPresentUnknown
)

type FiveGProSePC5QoSParametersExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      FiveGProSePC5QoSParametersExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSePC5QoSParametersExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveG-ProSePC5QoSParameters-ExtIEsEntry", v)
}

// DecodeFiveGProSePC5QoSParametersExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSePC5QoSParametersExtIEsEntry(data []byte) (FiveGProSePC5QoSParametersExtIEsEntry, int, error) {
	var v FiveGProSePC5QoSParametersExtIEsEntry
	err := ngapcodec.DecodeParameter("FiveG-ProSePC5QoSParameters-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSePC5QoSParametersExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
