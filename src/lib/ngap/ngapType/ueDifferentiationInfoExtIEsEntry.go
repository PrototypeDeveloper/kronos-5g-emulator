// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEDifferentiationInfoExtIEsEntryPresent uint32

const (
	UEDifferentiationInfoExtIEsEntryPresentNothing UEDifferentiationInfoExtIEsEntryPresent = iota
	UEDifferentiationInfoExtIEsEntryPresentUnknown
)

type UEDifferentiationInfoExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEDifferentiationInfoExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEDifferentiationInfoExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-DifferentiationInfo-ExtIEsEntry", v)
}

// DecodeUEDifferentiationInfoExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEDifferentiationInfoExtIEsEntry(data []byte) (UEDifferentiationInfoExtIEsEntry, int, error) {
	var v UEDifferentiationInfoExtIEsEntry
	err := ngapcodec.DecodeParameter("UE-DifferentiationInfo-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEDifferentiationInfoExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
