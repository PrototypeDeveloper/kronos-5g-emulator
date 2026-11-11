// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UENGAPIDPairExtIEsEntryPresent uint32

const (
	UENGAPIDPairExtIEsEntryPresentNothing UENGAPIDPairExtIEsEntryPresent = iota
	UENGAPIDPairExtIEsEntryPresentUnknown
)

type UENGAPIDPairExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UENGAPIDPairExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UENGAPIDPairExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-NGAP-ID-pair-ExtIEsEntry", v)
}

// DecodeUENGAPIDPairExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUENGAPIDPairExtIEsEntry(data []byte) (UENGAPIDPairExtIEsEntry, int, error) {
	var v UENGAPIDPairExtIEsEntry
	err := ngapcodec.DecodeParameter("UE-NGAP-ID-pair-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UENGAPIDPairExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
