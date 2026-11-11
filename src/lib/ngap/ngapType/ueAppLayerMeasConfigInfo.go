// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEAppLayerMeasConfigInfo represents ASN.1 UEAppLayerMeasConfigInfo.
type UEAppLayerMeasConfigInfo struct {
	QoEReference                      QoEReference
	ServiceType                       ServiceType
	HasAreaScopeOfQMC                 bool
	AreaScopeOfQMC                    AreaScopeOfQMC
	MeasCollEntityIPAddress           TransportLayerAddress
	HasQoEMeasurementStatus           bool
	QoEMeasurementStatus              UEAppLayerMeasConfigInfoQoEMeasurementStatus `aper:"valueExt,valueLB:0,valueUB:0"`
	HasContainerForAppLayerMeasConfig bool
	ContainerForAppLayerMeasConfig    []byte `aper:"sizeLB:1,sizeUB:8000"`
	HasMeasConfigAppLayerID           bool
	MeasConfigAppLayerID              int64 `aper:"valueExt,valueLB:0,valueUB:15"`
	HasSliceSupportListQMC            bool
	SliceSupportListQMC               SliceSupportListQMC
	HasMDTAlignmentInfo               bool
	MDTAlignmentInfo                  MDTAlignmentInfo
	HasAvailableRANVisibleQoEMetrics  bool
	AvailableRANVisibleQoEMetrics     AvailableRANVisibleQoEMetrics
	HasIEExtensions                   bool
	IEExtensions                      []UEAppLayerMeasConfigInfoExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAppLayerMeasConfigInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEAppLayerMeasConfigInfo", v)
}

// DecodeUEAppLayerMeasConfigInfo decodes a standalone APER value and reports the consumed octets.
func DecodeUEAppLayerMeasConfigInfo(data []byte) (UEAppLayerMeasConfigInfo, int, error) {
	var v UEAppLayerMeasConfigInfo
	err := ngapcodec.DecodeParameter("UEAppLayerMeasConfigInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAppLayerMeasConfigInfo{}, 0, err
	}
	return v, len(data), err
}
