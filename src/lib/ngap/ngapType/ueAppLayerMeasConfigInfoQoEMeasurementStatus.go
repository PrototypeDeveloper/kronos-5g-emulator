// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEAppLayerMeasConfigInfoQoEMeasurementStatus represents ASN.1 UEAppLayerMeasConfigInfoQoEMeasurementStatus.
const (
	UEAppLayerMeasConfigInfoQoEMeasurementStatusPresentOngoing Enumerated = 0
)

type UEAppLayerMeasConfigInfoQoEMeasurementStatus struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v UEAppLayerMeasConfigInfoQoEMeasurementStatus) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:UEAppLayerMeasConfigInfoQoEMeasurementStatus", v)
}

// DecodeUEAppLayerMeasConfigInfoQoEMeasurementStatus decodes a standalone APER value and reports consumed octets.
func DecodeUEAppLayerMeasConfigInfoQoEMeasurementStatus(data []byte) (UEAppLayerMeasConfigInfoQoEMeasurementStatus, int, error) {
	var v UEAppLayerMeasConfigInfoQoEMeasurementStatus
	err := ngapcodec.DecodeParameter("inline:UEAppLayerMeasConfigInfoQoEMeasurementStatus", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAppLayerMeasConfigInfoQoEMeasurementStatus{}, 0, err
	}
	return v, len(data), err
}
