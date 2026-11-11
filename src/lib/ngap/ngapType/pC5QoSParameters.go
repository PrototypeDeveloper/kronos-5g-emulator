// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PC5QoSParameters represents ASN.1 PC5QoSParameters.
type PC5QoSParameters struct {
	Pc5QoSFlowList              PC5QoSFlowList
	HasPc5LinkAggregateBitRates bool
	Pc5LinkAggregateBitRates    BitRate
	HasIEExtensions             bool
	IEExtensions                []PC5QoSParametersExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PC5QoSParameters) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PC5QoSParameters", v)
}

// DecodePC5QoSParameters decodes a standalone APER value and reports the consumed octets.
func DecodePC5QoSParameters(data []byte) (PC5QoSParameters, int, error) {
	var v PC5QoSParameters
	err := ngapcodec.DecodeParameter("PC5QoSParameters", data, &v)
	if ngapcodec.IsFatal(err) {
		return PC5QoSParameters{}, 0, err
	}
	return v, len(data), err
}
