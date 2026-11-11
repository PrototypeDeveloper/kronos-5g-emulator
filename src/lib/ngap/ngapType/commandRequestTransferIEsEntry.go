// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CommandRequestTransferIEsEntryPresent uint32

const (
	CommandRequestTransferIEsEntryPresentNothing CommandRequestTransferIEsEntryPresent = iota
	CommandRequestTransferIEsEntryPresentAIoTCorrelationIdentifier
	CommandRequestTransferIEsEntryPresentRANAIOTDeviceNGAPID
	CommandRequestTransferIEsEntryPresentAIoTNASPDU
	CommandRequestTransferIEsEntryPresentAIoTCommandAssistanceInformation
	CommandRequestTransferIEsEntryPresentUnknown
)

type CommandRequestTransferIEsEntry struct {
	ID                               ProtocolIEID
	Criticality                      Criticality
	Present                          CommandRequestTransferIEsEntryPresent
	AIoTCorrelationIdentifier        AIoTCorrelationIdentifier        `ngap:"id:466,criticality:reject,presence:mandatory"`
	RANAIOTDeviceNGAPID              RANAIOTDeviceNGAPID              `ngap:"id:472,criticality:reject,presence:mandatory"`
	AIoTNASPDU                       AIoTNASPDU                       `ngap:"id:474,criticality:reject,presence:mandatory"`
	AIoTCommandAssistanceInformation AIoTCommandAssistanceInformation `ngap:"id:473,criticality:reject,presence:mandatory"`
	UnknownValue                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CommandRequestTransferIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CommandRequestTransferIEsEntry", v)
}

// DecodeCommandRequestTransferIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCommandRequestTransferIEsEntry(data []byte) (CommandRequestTransferIEsEntry, int, error) {
	var v CommandRequestTransferIEsEntry
	err := ngapcodec.DecodeParameter("CommandRequestTransferIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CommandRequestTransferIEsEntry{}, 0, err
	}
	return v, len(data), err
}
