// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSQoSFlowsToBeSetupItemExtIEsEntryPresent uint32

const (
	MBSQoSFlowsToBeSetupItemExtIEsEntryPresentNothing MBSQoSFlowsToBeSetupItemExtIEsEntryPresent = iota
	MBSQoSFlowsToBeSetupItemExtIEsEntryPresentUnknown
)

type MBSQoSFlowsToBeSetupItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSQoSFlowsToBeSetupItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSQoSFlowsToBeSetupItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-QoSFlowsToBeSetupItem-ExtIEsEntry", v)
}

// DecodeMBSQoSFlowsToBeSetupItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSQoSFlowsToBeSetupItemExtIEsEntry(data []byte) (MBSQoSFlowsToBeSetupItemExtIEsEntry, int, error) {
	var v MBSQoSFlowsToBeSetupItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-QoSFlowsToBeSetupItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSQoSFlowsToBeSetupItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
