// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NGAPIESupportInformationRequestItemExtIEsEntryPresent uint32

const (
	NGAPIESupportInformationRequestItemExtIEsEntryPresentNothing NGAPIESupportInformationRequestItemExtIEsEntryPresent = iota
	NGAPIESupportInformationRequestItemExtIEsEntryPresentUnknown
)

type NGAPIESupportInformationRequestItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NGAPIESupportInformationRequestItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGAPIESupportInformationRequestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGAPIESupportInformationRequestItem-ExtIEsEntry", v)
}

// DecodeNGAPIESupportInformationRequestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNGAPIESupportInformationRequestItemExtIEsEntry(data []byte) (NGAPIESupportInformationRequestItemExtIEsEntry, int, error) {
	var v NGAPIESupportInformationRequestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("NGAPIESupportInformationRequestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPIESupportInformationRequestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
