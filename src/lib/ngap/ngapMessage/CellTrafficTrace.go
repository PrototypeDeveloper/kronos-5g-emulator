// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// CellTrafficTrace represents ASN.1 CellTrafficTrace.
type CellTrafficTrace struct {
	AMFUENGAPID                    ngapType.AMFUENGAPID           `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                    ngapType.RANUENGAPID           `ngap:"id:85,criticality:reject,presence:mandatory"`
	NGRANTraceID                   ngapType.NGRANTraceID          `ngap:"id:44,criticality:ignore,presence:mandatory"`
	NGRANCGI                       ngapType.NGRANCGI              `ngap:"id:43,criticality:ignore,presence:mandatory"`
	TraceCollectionEntityIPAddress ngapType.TransportLayerAddress `ngap:"id:109,criticality:ignore,presence:mandatory"`
	HasPrivacyIndicator            bool
	PrivacyIndicator               ngapType.PrivacyIndicator `ngap:"id:256,criticality:ignore,presence:optional"`
	HasTraceCollectionEntityURI    bool
	TraceCollectionEntityURI       ngapType.URIAddress `ngap:"id:257,criticality:ignore,presence:optional"`
	AdditionalIEs                  []ngapType.UnknownIE
	UnknownExtensions              []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *CellTrafficTrace) Encode() ([]byte, error) { return ngapcodec.Encode("CellTrafficTrace", m) }

// DecodeCellTrafficTrace decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeCellTrafficTrace(data []byte) (*CellTrafficTrace, error) {
	var m CellTrafficTrace
	err := ngapcodec.Decode("CellTrafficTrace", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
