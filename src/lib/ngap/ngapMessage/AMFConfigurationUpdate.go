// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// AMFConfigurationUpdate represents ASN.1 AMFConfigurationUpdate.
type AMFConfigurationUpdate struct {
	HasAMFName                       bool
	AMFName                          ngapType.AMFName `ngap:"id:1,criticality:reject,presence:optional"`
	HasServedGUAMIList               bool
	ServedGUAMIList                  ngapType.ServedGUAMIList `ngap:"id:96,criticality:reject,presence:optional"`
	HasRelativeAMFCapacity           bool
	RelativeAMFCapacity              ngapType.RelativeAMFCapacity `ngap:"id:86,criticality:ignore,presence:optional"`
	HasPLMNSupportList               bool
	PLMNSupportList                  ngapType.PLMNSupportList `ngap:"id:80,criticality:reject,presence:optional"`
	HasAMFTNLAssociationToAddList    bool
	AMFTNLAssociationToAddList       ngapType.AMFTNLAssociationToAddList `ngap:"id:6,criticality:ignore,presence:optional"`
	HasAMFTNLAssociationToRemoveList bool
	AMFTNLAssociationToRemoveList    ngapType.AMFTNLAssociationToRemoveList `ngap:"id:7,criticality:ignore,presence:optional"`
	HasAMFTNLAssociationToUpdateList bool
	AMFTNLAssociationToUpdateList    ngapType.AMFTNLAssociationToUpdateList `ngap:"id:8,criticality:ignore,presence:optional"`
	HasExtendedAMFName               bool
	ExtendedAMFName                  ngapType.ExtendedAMFName `ngap:"id:274,criticality:ignore,presence:optional"`
	AdditionalIEs                    []ngapType.UnknownIE
	UnknownExtensions                []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *AMFConfigurationUpdate) Encode() ([]byte, error) {
	return ngapcodec.Encode("AMFConfigurationUpdate", m)
}

// DecodeAMFConfigurationUpdate decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeAMFConfigurationUpdate(data []byte) (*AMFConfigurationUpdate, error) {
	var m AMFConfigurationUpdate
	err := ngapcodec.Decode("AMFConfigurationUpdate", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
