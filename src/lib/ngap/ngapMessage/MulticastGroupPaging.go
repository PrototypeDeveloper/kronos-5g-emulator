// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastGroupPaging represents ASN.1 MulticastGroupPaging.
type MulticastGroupPaging struct {
	MBSSessionID                 ngapType.MBSSessionID `ngap:"id:299,criticality:ignore,presence:mandatory"`
	HasMBSServiceArea            bool
	MBSServiceArea               ngapType.MBSServiceArea               `ngap:"id:298,criticality:ignore,presence:optional"`
	MulticastGroupPagingAreaList ngapType.MulticastGroupPagingAreaList `ngap:"id:307,criticality:ignore,presence:mandatory"`
	AdditionalIEs                []ngapType.UnknownIE
	UnknownExtensions            []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastGroupPaging) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastGroupPaging", m)
}

// DecodeMulticastGroupPaging decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastGroupPaging(data []byte) (*MulticastGroupPaging, error) {
	var m MulticastGroupPaging
	err := ngapcodec.Decode("MulticastGroupPaging", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
