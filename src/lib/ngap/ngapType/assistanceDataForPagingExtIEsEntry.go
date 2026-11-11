// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AssistanceDataForPagingExtIEsEntryPresent uint32

const (
	AssistanceDataForPagingExtIEsEntryPresentNothing AssistanceDataForPagingExtIEsEntryPresent = iota
	AssistanceDataForPagingExtIEsEntryPresentNPNPagingAssistanceInformation
	AssistanceDataForPagingExtIEsEntryPresentPagingAssisDataforCEcapabUE
	AssistanceDataForPagingExtIEsEntryPresentUnknown
)

type AssistanceDataForPagingExtIEsEntry struct {
	ID                             ProtocolExtensionID
	Criticality                    Criticality
	Present                        AssistanceDataForPagingExtIEsEntryPresent
	NPNPagingAssistanceInformation NPNPagingAssistanceInformation `ngap:"id:260,criticality:ignore,presence:optional"`
	PagingAssisDataforCEcapabUE    PagingAssisDataforCEcapabUE    `ngap:"id:207,criticality:ignore,presence:optional"`
	UnknownValue                   []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssistanceDataForPagingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssistanceDataForPaging-ExtIEsEntry", v)
}

// DecodeAssistanceDataForPagingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAssistanceDataForPagingExtIEsEntry(data []byte) (AssistanceDataForPagingExtIEsEntry, int, error) {
	var v AssistanceDataForPagingExtIEsEntry
	err := ngapcodec.DecodeParameter("AssistanceDataForPaging-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssistanceDataForPagingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
