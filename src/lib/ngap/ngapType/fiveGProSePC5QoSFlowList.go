// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSePC5QoSFlowList represents ASN.1 FiveGProSePC5QoSFlowList.
type FiveGProSePC5QoSFlowList struct {
	List []FiveGProSePC5QoSFlowItem `aper:"sizeLB:1,sizeUB:2048"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSePC5QoSFlowList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSePC5QoSFlowList", v)
}

// DecodeFiveGProSePC5QoSFlowList decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSePC5QoSFlowList(data []byte) (FiveGProSePC5QoSFlowList, int, error) {
	var v FiveGProSePC5QoSFlowList
	err := ngapcodec.DecodeParameter("FiveGProSePC5QoSFlowList", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSePC5QoSFlowList{}, 0, err
	}
	return v, len(data), err
}
