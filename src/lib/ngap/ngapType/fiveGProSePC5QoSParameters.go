// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSePC5QoSParameters represents ASN.1 FiveG-ProSePC5QoSParameters.
type FiveGProSePC5QoSParameters struct {
	FiveGProSepc5QoSFlowList              FiveGProSePC5QoSFlowList
	HasFiveGProSepc5LinkAggregateBitRates bool
	FiveGProSepc5LinkAggregateBitRates    BitRate
	HasIEExtensions                       bool
	IEExtensions                          []FiveGProSePC5QoSParametersExtIEsEntry
	UnknownExtensions                     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSePC5QoSParameters) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveG-ProSePC5QoSParameters", v)
}

// DecodeFiveGProSePC5QoSParameters decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSePC5QoSParameters(data []byte) (FiveGProSePC5QoSParameters, int, error) {
	var v FiveGProSePC5QoSParameters
	err := ngapcodec.DecodeParameter("FiveG-ProSePC5QoSParameters", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSePC5QoSParameters{}, 0, err
	}
	return v, len(data), err
}
