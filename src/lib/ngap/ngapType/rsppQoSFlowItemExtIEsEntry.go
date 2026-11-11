// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RSPPQoSFlowItemExtIEsEntryPresent uint32

const (
	RSPPQoSFlowItemExtIEsEntryPresentNothing RSPPQoSFlowItemExtIEsEntryPresent = iota
	RSPPQoSFlowItemExtIEsEntryPresentUnknown
)

type RSPPQoSFlowItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RSPPQoSFlowItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RSPPQoSFlowItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RSPPQoSFlowItem-ExtIEsEntry", v)
}

// DecodeRSPPQoSFlowItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRSPPQoSFlowItemExtIEsEntry(data []byte) (RSPPQoSFlowItemExtIEsEntry, int, error) {
	var v RSPPQoSFlowItemExtIEsEntry
	err := ngapcodec.DecodeParameter("RSPPQoSFlowItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RSPPQoSFlowItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
