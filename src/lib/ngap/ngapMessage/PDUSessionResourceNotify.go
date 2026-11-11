// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceNotify represents ASN.1 PDUSessionResourceNotify.
type PDUSessionResourceNotify struct {
	AMFUENGAPID                          ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                          ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasPDUSessionResourceNotifyList      bool
	PDUSessionResourceNotifyList         ngapType.PDUSessionResourceNotifyList `ngap:"id:66,criticality:reject,presence:optional"`
	HasPDUSessionResourceReleasedListNot bool
	PDUSessionResourceReleasedListNot    ngapType.PDUSessionResourceReleasedListNot `ngap:"id:67,criticality:ignore,presence:optional"`
	HasUserLocationInformation           bool
	UserLocationInformation              ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	AdditionalIEs                        []ngapType.UnknownIE
	UnknownExtensions                    []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceNotify) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceNotify", m)
}

// DecodePDUSessionResourceNotify decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceNotify(data []byte) (*PDUSessionResourceNotify, error) {
	var m PDUSessionResourceNotify
	err := ngapcodec.Decode("PDUSessionResourceNotify", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
