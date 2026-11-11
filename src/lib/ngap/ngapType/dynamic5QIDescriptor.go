// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// Dynamic5QIDescriptor represents ASN.1 Dynamic5QIDescriptor.
type Dynamic5QIDescriptor struct {
	PriorityLevelQos          PriorityLevelQos
	PacketDelayBudget         PacketDelayBudget
	PacketErrorRate           PacketErrorRate
	HasFiveQI                 bool
	FiveQI                    FiveQI
	HasDelayCritical          bool
	DelayCritical             DelayCritical
	HasAveragingWindow        bool
	AveragingWindow           AveragingWindow
	HasMaximumDataBurstVolume bool
	MaximumDataBurstVolume    MaximumDataBurstVolume
	HasIEExtensions           bool
	IEExtensions              []Dynamic5QIDescriptorExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v Dynamic5QIDescriptor) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Dynamic5QIDescriptor", v)
}

// DecodeDynamic5QIDescriptor decodes a standalone APER value and reports the consumed octets.
func DecodeDynamic5QIDescriptor(data []byte) (Dynamic5QIDescriptor, int, error) {
	var v Dynamic5QIDescriptor
	err := ngapcodec.DecodeParameter("Dynamic5QIDescriptor", data, &v)
	if ngapcodec.IsFatal(err) {
		return Dynamic5QIDescriptor{}, 0, err
	}
	return v, len(data), err
}
