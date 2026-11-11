// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NonDynamic5QIDescriptor represents ASN.1 NonDynamic5QIDescriptor.
type NonDynamic5QIDescriptor struct {
	FiveQI                    FiveQI
	HasPriorityLevelQos       bool
	PriorityLevelQos          PriorityLevelQos
	HasAveragingWindow        bool
	AveragingWindow           AveragingWindow
	HasMaximumDataBurstVolume bool
	MaximumDataBurstVolume    MaximumDataBurstVolume
	HasIEExtensions           bool
	IEExtensions              []NonDynamic5QIDescriptorExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NonDynamic5QIDescriptor) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NonDynamic5QIDescriptor", v)
}

// DecodeNonDynamic5QIDescriptor decodes a standalone APER value and reports the consumed octets.
func DecodeNonDynamic5QIDescriptor(data []byte) (NonDynamic5QIDescriptor, int, error) {
	var v NonDynamic5QIDescriptor
	err := ngapcodec.DecodeParameter("NonDynamic5QIDescriptor", data, &v)
	if ngapcodec.IsFatal(err) {
		return NonDynamic5QIDescriptor{}, 0, err
	}
	return v, len(data), err
}
