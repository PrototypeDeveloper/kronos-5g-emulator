// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEDifferentiationInfo represents ASN.1 UE-DifferentiationInfo.
type UEDifferentiationInfo struct {
	HasPeriodicCommunicationIndicator bool
	PeriodicCommunicationIndicator    UEDifferentiationInfoPeriodicCommunicationIndicator `aper:"valueExt,valueLB:0,valueUB:1"`
	HasPeriodicTime                   bool
	PeriodicTime                      int64 `aper:"valueExt,valueLB:1,valueUB:3600"`
	HasScheduledCommunicationTime     bool
	ScheduledCommunicationTime        ScheduledCommunicationTime
	HasStationaryIndication           bool
	StationaryIndication              UEDifferentiationInfoStationaryIndication `aper:"valueExt,valueLB:0,valueUB:1"`
	HasTrafficProfile                 bool
	TrafficProfile                    UEDifferentiationInfoTrafficProfile `aper:"valueExt,valueLB:0,valueUB:2"`
	HasBatteryIndication              bool
	BatteryIndication                 UEDifferentiationInfoBatteryIndication `aper:"valueExt,valueLB:0,valueUB:2"`
	HasIEExtensions                   bool
	IEExtensions                      []UEDifferentiationInfoExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEDifferentiationInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-DifferentiationInfo", v)
}

// DecodeUEDifferentiationInfo decodes a standalone APER value and reports the consumed octets.
func DecodeUEDifferentiationInfo(data []byte) (UEDifferentiationInfo, int, error) {
	var v UEDifferentiationInfo
	err := ngapcodec.DecodeParameter("UE-DifferentiationInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEDifferentiationInfo{}, 0, err
	}
	return v, len(data), err
}
