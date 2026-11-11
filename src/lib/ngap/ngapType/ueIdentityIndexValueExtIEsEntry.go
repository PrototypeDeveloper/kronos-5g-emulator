// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEIdentityIndexValueExtIEsEntryPresent uint32

const (
	UEIdentityIndexValueExtIEsEntryPresentNothing UEIdentityIndexValueExtIEsEntryPresent = iota
	UEIdentityIndexValueExtIEsEntryPresentUnknown
)

type UEIdentityIndexValueExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      UEIdentityIndexValueExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEIdentityIndexValueExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEIdentityIndexValue-ExtIEsEntry", v)
}

// DecodeUEIdentityIndexValueExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEIdentityIndexValueExtIEsEntry(data []byte) (UEIdentityIndexValueExtIEsEntry, int, error) {
	var v UEIdentityIndexValueExtIEsEntry
	err := ngapcodec.DecodeParameter("UEIdentityIndexValue-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEIdentityIndexValueExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
