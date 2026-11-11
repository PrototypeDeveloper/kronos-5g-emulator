// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeSyncAssistanceInfo represents ASN.1 TimeSyncAssistanceInfo.
type TimeSyncAssistanceInfo struct {
	TimeDistributionIndication TimeSyncAssistanceInfoTimeDistributionIndication `aper:"valueExt,valueLB:0,valueUB:1"`
	HasUUTimeSyncErrorBudget   bool
	UUTimeSyncErrorBudget      int64 `aper:"valueExt,valueLB:1,valueUB:1000000"`
	HasIEExtensions            bool
	IEExtensions               []TimeSyncAssistanceInfoExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeSyncAssistanceInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TimeSyncAssistanceInfo", v)
}

// DecodeTimeSyncAssistanceInfo decodes a standalone APER value and reports the consumed octets.
func DecodeTimeSyncAssistanceInfo(data []byte) (TimeSyncAssistanceInfo, int, error) {
	var v TimeSyncAssistanceInfo
	err := ngapcodec.DecodeParameter("TimeSyncAssistanceInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeSyncAssistanceInfo{}, 0, err
	}
	return v, len(data), err
}
