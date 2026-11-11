// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosCharacteristics represents ASN.1 QosCharacteristics.
type QosCharacteristicsPresent uint32

const (
	QosCharacteristicsPresentNothing QosCharacteristicsPresent = iota
	QosCharacteristicsPresentNonDynamic5QI
	QosCharacteristicsPresentDynamic5QI
	QosCharacteristicsPresentChoiceExtensions
)

type QosCharacteristics struct {
	Present          QosCharacteristicsPresent
	NonDynamic5QI    NonDynamic5QIDescriptor
	Dynamic5QI       Dynamic5QIDescriptor
	ChoiceExtensions QosCharacteristicsExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosCharacteristics) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosCharacteristics", v)
}

// DecodeQosCharacteristics decodes a standalone APER value and reports the consumed octets.
func DecodeQosCharacteristics(data []byte) (QosCharacteristics, int, error) {
	var v QosCharacteristics
	err := ngapcodec.DecodeParameter("QosCharacteristics", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosCharacteristics{}, 0, err
	}
	return v, len(data), err
}
