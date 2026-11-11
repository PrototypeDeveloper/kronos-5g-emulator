// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type N6JitterInformationExtIEsEntryPresent uint32

const (
	N6JitterInformationExtIEsEntryPresentNothing N6JitterInformationExtIEsEntryPresent = iota
	N6JitterInformationExtIEsEntryPresentUnknown
)

type N6JitterInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      N6JitterInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v N6JitterInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("N6JitterInformation-ExtIEsEntry", v)
}

// DecodeN6JitterInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeN6JitterInformationExtIEsEntry(data []byte) (N6JitterInformationExtIEsEntry, int, error) {
	var v N6JitterInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("N6JitterInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return N6JitterInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
