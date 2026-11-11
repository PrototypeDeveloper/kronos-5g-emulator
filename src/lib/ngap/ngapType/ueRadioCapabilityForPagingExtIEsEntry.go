// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UERadioCapabilityForPagingExtIEsEntryPresent uint32

const (
	UERadioCapabilityForPagingExtIEsEntryPresentNothing UERadioCapabilityForPagingExtIEsEntryPresent = iota
	UERadioCapabilityForPagingExtIEsEntryPresentUERadioCapabilityForPagingOfNBIoT
	UERadioCapabilityForPagingExtIEsEntryPresentUnknown
)

type UERadioCapabilityForPagingExtIEsEntry struct {
	ID                                ProtocolExtensionID
	Criticality                       Criticality
	Present                           UERadioCapabilityForPagingExtIEsEntryPresent
	UERadioCapabilityForPagingOfNBIoT UERadioCapabilityForPagingOfNBIoT `ngap:"id:214,criticality:ignore,presence:optional"`
	UnknownValue                      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERadioCapabilityForPagingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERadioCapabilityForPaging-ExtIEsEntry", v)
}

// DecodeUERadioCapabilityForPagingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUERadioCapabilityForPagingExtIEsEntry(data []byte) (UERadioCapabilityForPagingExtIEsEntry, int, error) {
	var v UERadioCapabilityForPagingExtIEsEntry
	err := ngapcodec.DecodeParameter("UERadioCapabilityForPaging-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERadioCapabilityForPagingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
