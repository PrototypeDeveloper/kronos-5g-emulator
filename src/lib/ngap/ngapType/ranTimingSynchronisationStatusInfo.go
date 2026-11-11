// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RANTimingSynchronisationStatusInfo represents ASN.1 RANTimingSynchronisationStatusInfo.
type RANTimingSynchronisationStatusInfo struct {
	HasSynchronisationState    bool
	SynchronisationState       RANTimingSynchronisationStatusInfoSynchronisationState `aper:"valueExt,valueLB:0,valueUB:2"`
	HasTraceabletoUTC          bool
	TraceabletoUTC             RANTimingSynchronisationStatusInfoTraceabletoUTC `aper:"valueExt,valueLB:0,valueUB:1"`
	HasTraceabletoGNSS         bool
	TraceabletoGNSS            RANTimingSynchronisationStatusInfoTraceabletoGNSS `aper:"valueExt,valueLB:0,valueUB:1"`
	HasClockFrequencyStability bool
	ClockFrequencyStability    BitString `aper:"sizeLB:16,sizeUB:16"`
	HasClockAccuracy           bool
	ClockAccuracy              ClockAccuracy
	HasParentTImeSource        bool
	ParentTImeSource           ParentTImeSource
	HasIEExtensions            bool
	IEExtensions               []RANTimingSynchronisationStatusInfoExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RANTimingSynchronisationStatusInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RANTimingSynchronisationStatusInfo", v)
}

// DecodeRANTimingSynchronisationStatusInfo decodes a standalone APER value and reports the consumed octets.
func DecodeRANTimingSynchronisationStatusInfo(data []byte) (RANTimingSynchronisationStatusInfo, int, error) {
	var v RANTimingSynchronisationStatusInfo
	err := ngapcodec.DecodeParameter("RANTimingSynchronisationStatusInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return RANTimingSynchronisationStatusInfo{}, 0, err
	}
	return v, len(data), err
}
