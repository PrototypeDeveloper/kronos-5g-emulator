// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MobileIABMTUserLocationInformationExtIEsEntryPresent uint32

const (
	MobileIABMTUserLocationInformationExtIEsEntryPresentNothing MobileIABMTUserLocationInformationExtIEsEntryPresent = iota
	MobileIABMTUserLocationInformationExtIEsEntryPresentUnknown
)

type MobileIABMTUserLocationInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MobileIABMTUserLocationInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MobileIABMTUserLocationInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MobileIAB-MTUserLocationInformation-ExtIEsEntry", v)
}

// DecodeMobileIABMTUserLocationInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMobileIABMTUserLocationInformationExtIEsEntry(data []byte) (MobileIABMTUserLocationInformationExtIEsEntry, int, error) {
	var v MobileIABMTUserLocationInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("MobileIAB-MTUserLocationInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MobileIABMTUserLocationInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
