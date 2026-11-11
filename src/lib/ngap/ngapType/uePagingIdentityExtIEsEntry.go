// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEPagingIdentityExtIEsEntryPresent uint32

const (
	UEPagingIdentityExtIEsEntryPresentNothing UEPagingIdentityExtIEsEntryPresent = iota
	UEPagingIdentityExtIEsEntryPresentUnknown
)

type UEPagingIdentityExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      UEPagingIdentityExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPagingIdentityExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEPagingIdentity-ExtIEsEntry", v)
}

// DecodeUEPagingIdentityExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEPagingIdentityExtIEsEntry(data []byte) (UEPagingIdentityExtIEsEntry, int, error) {
	var v UEPagingIdentityExtIEsEntry
	err := ngapcodec.DecodeParameter("UEPagingIdentity-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPagingIdentityExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
