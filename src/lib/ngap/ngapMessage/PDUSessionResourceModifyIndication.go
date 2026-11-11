// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceModifyIndication represents ASN.1 PDUSessionResourceModifyIndication.
type PDUSessionResourceModifyIndication struct {
	AMFUENGAPID                        ngapType.AMFUENGAPID                        `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                        ngapType.RANUENGAPID                        `ngap:"id:85,criticality:reject,presence:mandatory"`
	PDUSessionResourceModifyListModInd ngapType.PDUSessionResourceModifyListModInd `ngap:"id:63,criticality:reject,presence:mandatory"`
	HasUserLocationInformation         bool
	UserLocationInformation            ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	AdditionalIEs                      []ngapType.UnknownIE
	UnknownExtensions                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceModifyIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceModifyIndication", m)
}

// DecodePDUSessionResourceModifyIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceModifyIndication(data []byte) (*PDUSessionResourceModifyIndication, error) {
	var m PDUSessionResourceModifyIndication
	err := ngapcodec.Decode("PDUSessionResourceModifyIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
