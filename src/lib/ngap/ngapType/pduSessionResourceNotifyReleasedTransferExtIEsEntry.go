// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionResourceNotifyReleasedTransferExtIEsEntryPresent uint32

const (
	PDUSessionResourceNotifyReleasedTransferExtIEsEntryPresentNothing PDUSessionResourceNotifyReleasedTransferExtIEsEntryPresent = iota
	PDUSessionResourceNotifyReleasedTransferExtIEsEntryPresentSecondaryRATUsageInformation
	PDUSessionResourceNotifyReleasedTransferExtIEsEntryPresentUserPlaneErrorIndicator
	PDUSessionResourceNotifyReleasedTransferExtIEsEntryPresentUnknown
)

type PDUSessionResourceNotifyReleasedTransferExtIEsEntry struct {
	ID                           ProtocolExtensionID
	Criticality                  Criticality
	Present                      PDUSessionResourceNotifyReleasedTransferExtIEsEntryPresent
	SecondaryRATUsageInformation SecondaryRATUsageInformation `ngap:"id:144,criticality:ignore,presence:optional"`
	UserPlaneErrorIndicator      UserPlaneErrorIndicator      `ngap:"id:429,criticality:ignore,presence:optional"`
	UnknownValue                 []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionResourceNotifyReleasedTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionResourceNotifyReleasedTransfer-ExtIEsEntry", v)
}

// DecodePDUSessionResourceNotifyReleasedTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionResourceNotifyReleasedTransferExtIEsEntry(data []byte) (PDUSessionResourceNotifyReleasedTransferExtIEsEntry, int, error) {
	var v PDUSessionResourceNotifyReleasedTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionResourceNotifyReleasedTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionResourceNotifyReleasedTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
