// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RANfeedbacktypeProactiveExtIEsEntryPresent uint32

const (
	RANfeedbacktypeProactiveExtIEsEntryPresentNothing RANfeedbacktypeProactiveExtIEsEntryPresent = iota
	RANfeedbacktypeProactiveExtIEsEntryPresentUnknown
)

type RANfeedbacktypeProactiveExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RANfeedbacktypeProactiveExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANfeedbacktypeProactiveExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANfeedbacktype-proactive-ExtIEsEntry", v)
}

// DecodeRANfeedbacktypeProactiveExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRANfeedbacktypeProactiveExtIEsEntry(data []byte) (RANfeedbacktypeProactiveExtIEsEntry, int, error) {
	var v RANfeedbacktypeProactiveExtIEsEntry
	err := ngapcodec.DecodeParameter("RANfeedbacktype-proactive-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANfeedbacktypeProactiveExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
