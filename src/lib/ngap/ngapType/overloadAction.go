// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// OverloadAction represents ASN.1 OverloadAction.
const (
	OverloadActionPresentRejectNonEmergencyMoDt                                    Enumerated = 0
	OverloadActionPresentRejectRrcCrSignalling                                     Enumerated = 1
	OverloadActionPresentPermitEmergencySessionsAndMobileTerminatedServicesOnly    Enumerated = 2
	OverloadActionPresentPermitHighPrioritySessionsAndMobileTerminatedServicesOnly Enumerated = 3
)

type OverloadAction struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v OverloadAction) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("OverloadAction", v)
}

// DecodeOverloadAction decodes a standalone APER value and reports the consumed octets.
func DecodeOverloadAction(data []byte) (OverloadAction, int, error) {
	var v OverloadAction
	err := ngapcodec.DecodeParameter("OverloadAction", data, &v)
	if ngapcodec.IsFatal(err) {
		return OverloadAction{}, 0, err
	}
	return v, len(data), err
}
