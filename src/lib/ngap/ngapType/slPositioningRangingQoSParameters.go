// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SLPositioningRangingQoSParameters represents ASN.1 SLPositioningRangingQoSParameters.
type SLPositioningRangingQoSParameters struct {
	RSPPQoSFlowList              RSPPQoSFlowList
	HasRSPPLinkAggregateBitRates bool
	RSPPLinkAggregateBitRates    BitRate
	HasIEExtensions              bool
	IEExtensions                 []SLPositioningRangingQoSParametersExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SLPositioningRangingQoSParameters) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SLPositioningRangingQoSParameters", v)
}

// DecodeSLPositioningRangingQoSParameters decodes a standalone APER value and reports the consumed octets.
func DecodeSLPositioningRangingQoSParameters(data []byte) (SLPositioningRangingQoSParameters, int, error) {
	var v SLPositioningRangingQoSParameters
	err := ngapcodec.DecodeParameter("SLPositioningRangingQoSParameters", data, &v)
	if ngapcodec.IsFatal(err) {
		return SLPositioningRangingQoSParameters{}, 0, err
	}
	return v, len(data), err
}
