// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// WriteReplaceWarningRequest represents ASN.1 WriteReplaceWarningRequest.
type WriteReplaceWarningRequest struct {
	MessageIdentifier              ngapType.MessageIdentifier `ngap:"id:35,criticality:reject,presence:mandatory"`
	SerialNumber                   ngapType.SerialNumber      `ngap:"id:95,criticality:reject,presence:mandatory"`
	HasWarningAreaList             bool
	WarningAreaList                ngapType.WarningAreaList             `ngap:"id:122,criticality:ignore,presence:optional"`
	RepetitionPeriod               ngapType.RepetitionPeriod            `ngap:"id:87,criticality:reject,presence:mandatory"`
	NumberOfBroadcastsRequested    ngapType.NumberOfBroadcastsRequested `ngap:"id:47,criticality:reject,presence:mandatory"`
	HasWarningType                 bool
	WarningType                    ngapType.WarningType `ngap:"id:125,criticality:ignore,presence:optional"`
	HasWarningSecurityInfo         bool
	WarningSecurityInfo            ngapType.WarningSecurityInfo `ngap:"id:124,criticality:ignore,presence:optional"`
	HasDataCodingScheme            bool
	DataCodingScheme               ngapType.DataCodingScheme `ngap:"id:20,criticality:ignore,presence:optional"`
	HasWarningMessageContents      bool
	WarningMessageContents         ngapType.WarningMessageContents `ngap:"id:123,criticality:ignore,presence:optional"`
	HasConcurrentWarningMessageInd bool
	ConcurrentWarningMessageInd    ngapType.ConcurrentWarningMessageInd `ngap:"id:17,criticality:reject,presence:optional"`
	HasWarningAreaCoordinates      bool
	WarningAreaCoordinates         ngapType.WarningAreaCoordinates `ngap:"id:141,criticality:ignore,presence:optional"`
	AdditionalIEs                  []ngapType.UnknownIE
	UnknownExtensions              []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *WriteReplaceWarningRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("WriteReplaceWarningRequest", m)
}

// DecodeWriteReplaceWarningRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeWriteReplaceWarningRequest(data []byte) (*WriteReplaceWarningRequest, error) {
	var m WriteReplaceWarningRequest
	err := ngapcodec.Decode("WriteReplaceWarningRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
