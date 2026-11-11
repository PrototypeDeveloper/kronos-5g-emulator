// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QoSFlowTSCItemExtIEsEntryPresent uint32

const (
	QoSFlowTSCItemExtIEsEntryPresentNothing QoSFlowTSCItemExtIEsEntryPresent = iota
	QoSFlowTSCItemExtIEsEntryPresentUnknown
)

type QoSFlowTSCItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QoSFlowTSCItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoSFlowTSCItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoSFlowTSCItem-ExtIEsEntry", v)
}

// DecodeQoSFlowTSCItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQoSFlowTSCItemExtIEsEntry(data []byte) (QoSFlowTSCItemExtIEsEntry, int, error) {
	var v QoSFlowTSCItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QoSFlowTSCItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoSFlowTSCItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
