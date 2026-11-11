// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// RetrieveUEInformation represents ASN.1 RetrieveUEInformation.
type RetrieveUEInformation struct {
	FiveGSTMSI        ngapType.FiveGSTMSI `ngap:"id:26,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *RetrieveUEInformation) Encode() ([]byte, error) {
	return ngapcodec.Encode("RetrieveUEInformation", m)
}

// DecodeRetrieveUEInformation decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeRetrieveUEInformation(data []byte) (*RetrieveUEInformation, error) {
	var m RetrieveUEInformation
	err := ngapcodec.Decode("RetrieveUEInformation", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
