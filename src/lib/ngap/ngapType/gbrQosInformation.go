// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GBRQosInformation represents ASN.1 GBR-QosInformation.
type GBRQosInformation struct {
	MaximumFlowBitRateDL       BitRate
	MaximumFlowBitRateUL       BitRate
	GuaranteedFlowBitRateDL    BitRate
	GuaranteedFlowBitRateUL    BitRate
	HasNotificationControl     bool
	NotificationControl        NotificationControl
	HasMaximumPacketLossRateDL bool
	MaximumPacketLossRateDL    PacketLossRate
	HasMaximumPacketLossRateUL bool
	MaximumPacketLossRateUL    PacketLossRate
	HasIEExtensions            bool
	IEExtensions               []GBRQosInformationExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GBRQosInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GBR-QosInformation", v)
}

// DecodeGBRQosInformation decodes a standalone APER value and reports the consumed octets.
func DecodeGBRQosInformation(data []byte) (GBRQosInformation, int, error) {
	var v GBRQosInformation
	err := ngapcodec.DecodeParameter("GBR-QosInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return GBRQosInformation{}, 0, err
	}
	return v, len(data), err
}
