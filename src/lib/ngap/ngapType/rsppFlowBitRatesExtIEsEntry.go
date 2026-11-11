// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RSPPFlowBitRatesExtIEsEntryPresent uint32

const (
	RSPPFlowBitRatesExtIEsEntryPresentNothing RSPPFlowBitRatesExtIEsEntryPresent = iota
	RSPPFlowBitRatesExtIEsEntryPresentUnknown
)

type RSPPFlowBitRatesExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RSPPFlowBitRatesExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RSPPFlowBitRatesExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RSPPFlowBitRates-ExtIEsEntry", v)
}

// DecodeRSPPFlowBitRatesExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRSPPFlowBitRatesExtIEsEntry(data []byte) (RSPPFlowBitRatesExtIEsEntry, int, error) {
	var v RSPPFlowBitRatesExtIEsEntry
	err := ngapcodec.DecodeParameter("RSPPFlowBitRates-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RSPPFlowBitRatesExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
