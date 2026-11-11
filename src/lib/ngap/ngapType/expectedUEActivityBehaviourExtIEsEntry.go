// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ExpectedUEActivityBehaviourExtIEsEntryPresent uint32

const (
	ExpectedUEActivityBehaviourExtIEsEntryPresentNothing ExpectedUEActivityBehaviourExtIEsEntryPresent = iota
	ExpectedUEActivityBehaviourExtIEsEntryPresentUnknown
)

type ExpectedUEActivityBehaviourExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ExpectedUEActivityBehaviourExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExpectedUEActivityBehaviourExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExpectedUEActivityBehaviour-ExtIEsEntry", v)
}

// DecodeExpectedUEActivityBehaviourExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeExpectedUEActivityBehaviourExtIEsEntry(data []byte) (ExpectedUEActivityBehaviourExtIEsEntry, int, error) {
	var v ExpectedUEActivityBehaviourExtIEsEntry
	err := ngapcodec.DecodeParameter("ExpectedUEActivityBehaviour-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExpectedUEActivityBehaviourExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
