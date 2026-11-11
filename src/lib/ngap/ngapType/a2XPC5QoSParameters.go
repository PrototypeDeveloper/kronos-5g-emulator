// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// A2XPC5QoSParameters represents ASN.1 A2X-PC5-QoS-Parameters.
type A2XPC5QoSParameters struct {
	A2XPC5QoSFlowList              A2XPC5QoSFlowList
	HasA2XPC5LinkAggregateBitRates bool
	A2XPC5LinkAggregateBitRates    BitRate
	HasIEExtensions                bool
	IEExtensions                   []A2XPC5QoSParametersExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v A2XPC5QoSParameters) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("A2X-PC5-QoS-Parameters", v)
}

// DecodeA2XPC5QoSParameters decodes a standalone APER value and reports the consumed octets.
func DecodeA2XPC5QoSParameters(data []byte) (A2XPC5QoSParameters, int, error) {
	var v A2XPC5QoSParameters
	err := ngapcodec.DecodeParameter("A2X-PC5-QoS-Parameters", data, &v)
	if ngapcodec.IsFatal(err) {
		return A2XPC5QoSParameters{}, 0, err
	}
	return v, len(data), err
}
