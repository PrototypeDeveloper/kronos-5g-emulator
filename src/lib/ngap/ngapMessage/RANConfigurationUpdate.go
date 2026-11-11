// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// RANConfigurationUpdate represents ASN.1 RANConfigurationUpdate.
type RANConfigurationUpdate struct {
	HasRANNodeName                     bool
	RANNodeName                        ngapType.RANNodeName `ngap:"id:82,criticality:ignore,presence:optional"`
	HasSupportedTAList                 bool
	SupportedTAList                    ngapType.SupportedTAList `ngap:"id:102,criticality:reject,presence:optional"`
	HasDefaultPagingDRX                bool
	DefaultPagingDRX                   ngapType.PagingDRX `ngap:"id:21,criticality:ignore,presence:optional"`
	HasGlobalRANNodeID                 bool
	GlobalRANNodeID                    ngapType.GlobalRANNodeID `ngap:"id:27,criticality:ignore,presence:optional"`
	HasNGRANTNLAssociationToRemoveList bool
	NGRANTNLAssociationToRemoveList    ngapType.NGRANTNLAssociationToRemoveList `ngap:"id:167,criticality:reject,presence:optional"`
	HasNBIoTDefaultPagingDRX           bool
	NBIoTDefaultPagingDRX              ngapType.NBIoTDefaultPagingDRX `ngap:"id:204,criticality:ignore,presence:optional"`
	HasExtendedRANNodeName             bool
	ExtendedRANNodeName                ngapType.ExtendedRANNodeName `ngap:"id:273,criticality:ignore,presence:optional"`
	HasAdditionalULI                   bool
	AdditionalULI                      ngapType.AdditionalULI `ngap:"id:483,criticality:ignore,presence:optional"`
	HasAIoTSupport                     bool
	AIoTSupport                        ngapType.AIoTSupport `ngap:"id:475,criticality:reject,presence:optional"`
	AdditionalIEs                      []ngapType.UnknownIE
	UnknownExtensions                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *RANConfigurationUpdate) Encode() ([]byte, error) {
	return ngapcodec.Encode("RANConfigurationUpdate", m)
}

// DecodeRANConfigurationUpdate decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeRANConfigurationUpdate(data []byte) (*RANConfigurationUpdate, error) {
	var m RANConfigurationUpdate
	err := ngapcodec.Decode("RANConfigurationUpdate", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
