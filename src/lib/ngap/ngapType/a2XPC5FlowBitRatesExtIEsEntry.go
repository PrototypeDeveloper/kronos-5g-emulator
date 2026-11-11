// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type A2XPC5FlowBitRatesExtIEsEntryPresent uint32

const (
	A2XPC5FlowBitRatesExtIEsEntryPresentNothing A2XPC5FlowBitRatesExtIEsEntryPresent = iota
	A2XPC5FlowBitRatesExtIEsEntryPresentUnknown
)

type A2XPC5FlowBitRatesExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      A2XPC5FlowBitRatesExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v A2XPC5FlowBitRatesExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("A2X-PC5-FlowBitRates-ExtIEsEntry", v)
}

// DecodeA2XPC5FlowBitRatesExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeA2XPC5FlowBitRatesExtIEsEntry(data []byte) (A2XPC5FlowBitRatesExtIEsEntry, int, error) {
	var v A2XPC5FlowBitRatesExtIEsEntry
	err := ngapcodec.DecodeParameter("A2X-PC5-FlowBitRates-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return A2XPC5FlowBitRatesExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
