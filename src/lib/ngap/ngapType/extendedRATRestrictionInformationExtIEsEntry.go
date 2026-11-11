// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ExtendedRATRestrictionInformationExtIEsEntryPresent uint32

const (
	ExtendedRATRestrictionInformationExtIEsEntryPresentNothing ExtendedRATRestrictionInformationExtIEsEntryPresent = iota
	ExtendedRATRestrictionInformationExtIEsEntryPresentUnknown
)

type ExtendedRATRestrictionInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ExtendedRATRestrictionInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedRATRestrictionInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExtendedRATRestrictionInformation-ExtIEsEntry", v)
}

// DecodeExtendedRATRestrictionInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedRATRestrictionInformationExtIEsEntry(data []byte) (ExtendedRATRestrictionInformationExtIEsEntry, int, error) {
	var v ExtendedRATRestrictionInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("ExtendedRATRestrictionInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedRATRestrictionInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
