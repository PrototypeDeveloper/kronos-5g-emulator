// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RANTimingSynchronisationStatusInfoExtIEsEntryPresent uint32

const (
	RANTimingSynchronisationStatusInfoExtIEsEntryPresentNothing RANTimingSynchronisationStatusInfoExtIEsEntryPresent = iota
	RANTimingSynchronisationStatusInfoExtIEsEntryPresentUnknown
)

type RANTimingSynchronisationStatusInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RANTimingSynchronisationStatusInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANTimingSynchronisationStatusInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANTimingSynchronisationStatusInfo-ExtIEsEntry", v)
}

// DecodeRANTimingSynchronisationStatusInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRANTimingSynchronisationStatusInfoExtIEsEntry(data []byte) (RANTimingSynchronisationStatusInfoExtIEsEntry, int, error) {
	var v RANTimingSynchronisationStatusInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("RANTimingSynchronisationStatusInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTimingSynchronisationStatusInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
