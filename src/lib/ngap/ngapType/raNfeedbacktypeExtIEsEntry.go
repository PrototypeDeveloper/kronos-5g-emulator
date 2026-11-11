// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RANfeedbacktypeExtIEsEntryPresent uint32

const (
	RANfeedbacktypeExtIEsEntryPresentNothing RANfeedbacktypeExtIEsEntryPresent = iota
	RANfeedbacktypeExtIEsEntryPresentUnknown
)

type RANfeedbacktypeExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      RANfeedbacktypeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANfeedbacktypeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANfeedbacktype-ExtIEsEntry", v)
}

// DecodeRANfeedbacktypeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRANfeedbacktypeExtIEsEntry(data []byte) (RANfeedbacktypeExtIEsEntry, int, error) {
	var v RANfeedbacktypeExtIEsEntry
	err := ngapcodec.DecodeParameter("RANfeedbacktype-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANfeedbacktypeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
