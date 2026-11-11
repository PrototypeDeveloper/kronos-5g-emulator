// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// LocationReportingControl represents ASN.1 LocationReportingControl.
type LocationReportingControl struct {
	AMFUENGAPID                  ngapType.AMFUENGAPID                  `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                  ngapType.RANUENGAPID                  `ngap:"id:85,criticality:reject,presence:mandatory"`
	LocationReportingRequestType ngapType.LocationReportingRequestType `ngap:"id:33,criticality:ignore,presence:mandatory"`
	AdditionalIEs                []ngapType.UnknownIE
	UnknownExtensions            []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *LocationReportingControl) Encode() ([]byte, error) {
	return ngapcodec.Encode("LocationReportingControl", m)
}

// DecodeLocationReportingControl decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeLocationReportingControl(data []byte) (*LocationReportingControl, error) {
	var m LocationReportingControl
	err := ngapcodec.Decode("LocationReportingControl", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
