// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type RequestedTNLInfoExtIEsEntryPresent uint32

const (
	RequestedTNLInfoExtIEsEntryPresentNothing RequestedTNLInfoExtIEsEntryPresent = iota
	RequestedTNLInfoExtIEsEntryPresentUnknown
)

type RequestedTNLInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      RequestedTNLInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RequestedTNLInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RequestedTNLInfo-ExtIEsEntry", v)
}

// DecodeRequestedTNLInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeRequestedTNLInfoExtIEsEntry(data []byte) (RequestedTNLInfoExtIEsEntry, int, error) {
	var v RequestedTNLInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("RequestedTNLInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return RequestedTNLInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
