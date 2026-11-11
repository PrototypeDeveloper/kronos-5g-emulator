// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RANfeedbacktypeReactiveExtIEsEntryPresent uint32

const (
	RANfeedbacktypeReactiveExtIEsEntryPresentNothing RANfeedbacktypeReactiveExtIEsEntryPresent = iota
	RANfeedbacktypeReactiveExtIEsEntryPresentUnknown
)

type RANfeedbacktypeReactiveExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RANfeedbacktypeReactiveExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANfeedbacktypeReactiveExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANfeedbacktype-reactive-ExtIEsEntry", v)
}

// DecodeRANfeedbacktypeReactiveExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRANfeedbacktypeReactiveExtIEsEntry(data []byte) (RANfeedbacktypeReactiveExtIEsEntry, int, error) {
	var v RANfeedbacktypeReactiveExtIEsEntry
	err := ngapcodec.DecodeParameter("RANfeedbacktype-reactive-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANfeedbacktypeReactiveExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
