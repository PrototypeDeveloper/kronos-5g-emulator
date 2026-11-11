// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceModifyRequest represents ASN.1 PDUSessionResourceModifyRequest.
type PDUSessionResourceModifyRequest struct {
	AMFUENGAPID                        ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                        ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasRANPagingPriority               bool
	RANPagingPriority                  ngapType.RANPagingPriority                  `ngap:"id:83,criticality:ignore,presence:optional"`
	PDUSessionResourceModifyListModReq ngapType.PDUSessionResourceModifyListModReq `ngap:"id:64,criticality:reject,presence:mandatory"`
	AdditionalIEs                      []ngapType.UnknownIE
	UnknownExtensions                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceModifyRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceModifyRequest", m)
}

// DecodePDUSessionResourceModifyRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceModifyRequest(data []byte) (*PDUSessionResourceModifyRequest, error) {
	var m PDUSessionResourceModifyRequest
	err := ngapcodec.Decode("PDUSessionResourceModifyRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
