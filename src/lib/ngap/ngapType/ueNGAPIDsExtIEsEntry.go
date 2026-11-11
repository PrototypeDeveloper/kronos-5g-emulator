// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UENGAPIDsExtIEsEntryPresent uint32

const (
	UENGAPIDsExtIEsEntryPresentNothing UENGAPIDsExtIEsEntryPresent = iota
	UENGAPIDsExtIEsEntryPresentUnknown
)

type UENGAPIDsExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      UENGAPIDsExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UENGAPIDsExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-NGAP-IDs-ExtIEsEntry", v)
}

// DecodeUENGAPIDsExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUENGAPIDsExtIEsEntry(data []byte) (UENGAPIDsExtIEsEntry, int, error) {
	var v UENGAPIDsExtIEsEntry
	err := ngapcodec.DecodeParameter("UE-NGAP-IDs-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UENGAPIDsExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
