// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ExpectedUEBehaviourExtIEsEntryPresent uint32

const (
	ExpectedUEBehaviourExtIEsEntryPresentNothing ExpectedUEBehaviourExtIEsEntryPresent = iota
	ExpectedUEBehaviourExtIEsEntryPresentUnknown
)

type ExpectedUEBehaviourExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ExpectedUEBehaviourExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedUEBehaviourExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedUEBehaviour-ExtIEsEntry", v)
}

// DecodeExpectedUEBehaviourExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedUEBehaviourExtIEsEntry(data []byte) (ExpectedUEBehaviourExtIEsEntry, int, error) {
	var v ExpectedUEBehaviourExtIEsEntry
	err := ngapcodec.DecodeParameter("ExpectedUEBehaviour-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedUEBehaviourExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
