// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NGRANRadioResourceStatusExtIEsEntryPresent uint32

const (
	NGRANRadioResourceStatusExtIEsEntryPresentNothing NGRANRadioResourceStatusExtIEsEntryPresent = iota
	NGRANRadioResourceStatusExtIEsEntryPresentUnknown
)

type NGRANRadioResourceStatusExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NGRANRadioResourceStatusExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANRadioResourceStatusExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-RadioResourceStatus-ExtIEsEntry", v)
}

// DecodeNGRANRadioResourceStatusExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANRadioResourceStatusExtIEsEntry(data []byte) (NGRANRadioResourceStatusExtIEsEntry, int, error) {
	var v NGRANRadioResourceStatusExtIEsEntry
	err := ngapcodec.DecodeParameter("NGRAN-RadioResourceStatus-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANRadioResourceStatusExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
