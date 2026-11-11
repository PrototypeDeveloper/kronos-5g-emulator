// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceReleaseCommand represents ASN.1 PDUSessionResourceReleaseCommand.
type PDUSessionResourceReleaseCommand struct {
	AMFUENGAPID                           ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                           ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasRANPagingPriority                  bool
	RANPagingPriority                     ngapType.RANPagingPriority `ngap:"id:83,criticality:ignore,presence:optional"`
	HasNASPDU                             bool
	NASPDU                                ngapType.NASPDU                                `ngap:"id:38,criticality:ignore,presence:optional"`
	PDUSessionResourceToReleaseListRelCmd ngapType.PDUSessionResourceToReleaseListRelCmd `ngap:"id:79,criticality:reject,presence:mandatory"`
	AdditionalIEs                         []ngapType.UnknownIE
	UnknownExtensions                     []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceReleaseCommand) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceReleaseCommand", m)
}

// DecodePDUSessionResourceReleaseCommand decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceReleaseCommand(data []byte) (*PDUSessionResourceReleaseCommand, error) {
	var m PDUSessionResourceReleaseCommand
	err := ngapcodec.Decode("PDUSessionResourceReleaseCommand", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
