// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FiveGProSePC5FlowBitRatesExtIEsEntryPresent uint32

const (
	FiveGProSePC5FlowBitRatesExtIEsEntryPresentNothing FiveGProSePC5FlowBitRatesExtIEsEntryPresent = iota
	FiveGProSePC5FlowBitRatesExtIEsEntryPresentUnknown
)

type FiveGProSePC5FlowBitRatesExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      FiveGProSePC5FlowBitRatesExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSePC5FlowBitRatesExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSePC5FlowBitRates-ExtIEsEntry", v)
}

// DecodeFiveGProSePC5FlowBitRatesExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSePC5FlowBitRatesExtIEsEntry(data []byte) (FiveGProSePC5FlowBitRatesExtIEsEntry, int, error) {
	var v FiveGProSePC5FlowBitRatesExtIEsEntry
	err := ngapcodec.DecodeParameter("FiveGProSePC5FlowBitRates-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSePC5FlowBitRatesExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
