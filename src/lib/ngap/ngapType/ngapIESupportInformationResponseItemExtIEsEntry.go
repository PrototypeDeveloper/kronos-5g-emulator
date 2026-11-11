// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NGAPIESupportInformationResponseItemExtIEsEntryPresent uint32

const (
	NGAPIESupportInformationResponseItemExtIEsEntryPresentNothing NGAPIESupportInformationResponseItemExtIEsEntryPresent = iota
	NGAPIESupportInformationResponseItemExtIEsEntryPresentUnknown
)

type NGAPIESupportInformationResponseItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NGAPIESupportInformationResponseItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGAPIESupportInformationResponseItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGAPIESupportInformationResponseItem-ExtIEsEntry", v)
}

// DecodeNGAPIESupportInformationResponseItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNGAPIESupportInformationResponseItemExtIEsEntry(data []byte) (NGAPIESupportInformationResponseItemExtIEsEntry, int, error) {
	var v NGAPIESupportInformationResponseItemExtIEsEntry
	err := ngapcodec.DecodeParameter("NGAPIESupportInformationResponseItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGAPIESupportInformationResponseItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
