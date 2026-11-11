// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MeasurementThresholdL1LoggedMDT represents ASN.1 MeasurementThresholdL1LoggedMDT.
type MeasurementThresholdL1LoggedMDTPresent uint32

const (
	MeasurementThresholdL1LoggedMDTPresentNothing MeasurementThresholdL1LoggedMDTPresent = iota
	MeasurementThresholdL1LoggedMDTPresentThresholdRSRP
	MeasurementThresholdL1LoggedMDTPresentThresholdRSRQ
	MeasurementThresholdL1LoggedMDTPresentChoiceExtensions
)

type MeasurementThresholdL1LoggedMDT struct {
	Present          MeasurementThresholdL1LoggedMDTPresent
	ThresholdRSRP    ThresholdRSRP
	ThresholdRSRQ    ThresholdRSRQ
	ChoiceExtensions MeasurementThresholdL1LoggedMDTExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MeasurementThresholdL1LoggedMDT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MeasurementThresholdL1LoggedMDT", v)
}

// DecodeMeasurementThresholdL1LoggedMDT decodes a standalone APER value and reports the consumed octets.
func DecodeMeasurementThresholdL1LoggedMDT(data []byte) (MeasurementThresholdL1LoggedMDT, int, error) {
	var v MeasurementThresholdL1LoggedMDT
	err := ngapcodec.DecodeParameter("MeasurementThresholdL1LoggedMDT", data, &v)
	if ngapcodec.IsFatal(err) {
		return MeasurementThresholdL1LoggedMDT{}, 0, err
	}
	return v, len(data), err
}
