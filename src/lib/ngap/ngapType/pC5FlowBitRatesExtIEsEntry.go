// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PC5FlowBitRatesExtIEsEntryPresent uint32

const (
	PC5FlowBitRatesExtIEsEntryPresentNothing PC5FlowBitRatesExtIEsEntryPresent = iota
	PC5FlowBitRatesExtIEsEntryPresentUnknown
)

type PC5FlowBitRatesExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PC5FlowBitRatesExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PC5FlowBitRatesExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PC5FlowBitRates-ExtIEsEntry", v)
}

// DecodePC5FlowBitRatesExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePC5FlowBitRatesExtIEsEntry(data []byte) (PC5FlowBitRatesExtIEsEntry, int, error) {
	var v PC5FlowBitRatesExtIEsEntry
	err := ngapcodec.DecodeParameter("PC5FlowBitRates-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PC5FlowBitRatesExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
